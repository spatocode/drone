package opengl

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/spatocode/karid"
)

var _ karid.Renderer = (*glRenderer)(nil)

var vertexShaderSource = `
#version 330 core
layout (location = 0) in vec3 aPos;

void main() {
	gl_Position = vec4(aPos.x, aPos.y, aPos.z, 1.0);
}
` + "\x00"

var fragmentShaderSource = `
#version 330 core
out vec4 FragColor;

void main() {
	FragColor = vec4(1.0f, 0.5f, 0.2f, 1.0f);
}
` + "\x00"
)

// ShaderProgram represents an OpenGL program
type ShaderProgram uint32

// VertexBuffer represents an OpenGL vertex buffer object
type VertexBuffer uint32

// VertexArray represents an OpenGL vertex array object
type VertexArray uint32

type glRenderer struct {
	scene karid.Scene
	program ShaderProgram
	vao VertexArray
	vbo VertexBuffer
}

// Render renders a scene using OpenGL
func (glr *glRenderer) Render(scene karid.Scene) {
	glr.scene = scene
	vertexShader, err := glr.compileShaders(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}

	fragmentShader, err := glr.compileShaders(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	glr.linkShaders(vertexShader, fragmentShader)
	glr.createBuffers()
}

func (glr *glRenderer) genCoordVertices() []float32 {
	return []float32{
		-0.5, -0.5, 0.0, // left
		0.5, -0.5, 0.0, // right
		0.0, 0.5, 0.0, // top
	} 
}

// Draw draws the scene and it's objects
func (glr *glRenderer) Draw() {
	// make sure there are objects in the scene before drawing
	if glr.scene == nil {
		return
	}
	glr.clear()
	gl.UseProgram(uint32(glr.program))

	//gl.BindVertexArray(uint32(glr.vao)) // we have a single vao so no need binding it all the time but we do this to keep things simple
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(triangle)/3))
}

func (glr *glRenderer) linkShaders(vertex, fragment uint32) {
	var status int32
	shaderProgram := gl.CreateProgram()

	gl.AttachShader(shaderProgram, vertex)
	gl.AttachShader(shaderProgram, fragment)
	gl.LinkProgram(shaderProgram)

	gl.GetProgramiv(shaderProgram, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(shaderProgram, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(shaderProgram, logLength, nil, gl.Str(log))
		panic("failed to link shaders")
	}

	glr.program = ShaderProgram(shaderProgram)
	gl.DeleteShader(vertex)
	gl.DeleteShader(fragment)
}

func (glr *glRenderer) createBuffers() {
	vertices := glr.genCoordVertices()
	var vbo, vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.GenBuffers(1, &vbo)
	gl.BindVertexArray(vao)

	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(vertices), gl.Ptr(vertices), gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
	gl.EnableVertexAttribArray(0)

	// TODO: create a separate method for freeing buffers
	// unbind the buffer as the call to gl.VertexAttribPointer has registered the vbo
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	// gl.BindVertexArray(0) // you can unbind vao also

	glr.vao = VertexArray(vao)
	glr.vbo = VertexBuffer(vbo)
}

func (glr *glRenderer) Free() {
	//gl.DeteVertexArray(1, )
}

func (glr *glRenderer) compileShaders(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csource, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csource, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}
	
	return shader, nil
}

func (glr *glRenderer) clear() {
	gl.ClearColor(0, 0.5, 1.0, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

// New initializes OpenGL renderer
func New() karid.Renderer {
	err := gl.Init()
	if err != nil {
		karid.LogError("Failed to initialize OpenGL", err)
	}
	return &glRenderer{}
}
