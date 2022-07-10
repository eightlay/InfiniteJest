package shaders

const (
	// Basic vertex shader
	vertexShaderSource = `
		#version 330 core
		layout (location = 0) in vec3 aPos;
		layout (location = 1) in vec2 aTexCoord;
		
		out vec2 TexCoord;
		
		uniform mat4 model;
		uniform mat4 view;
		uniform mat4 projection;
		
		void main()
		{
			gl_Position = projection * view * model * vec4(aPos, 1.0);
			TexCoord = vec2(aTexCoord.x, aTexCoord.y);
		}
	` + "\x00"

	// Basic fragment shader
	fragmentShaderSource = `
		#version 330 core
		out vec4 FragColor;
		
		in vec2 TexCoord;
		
		uniform sampler2D ourTexture;
		
		void main()
		{
			FragColor = texture(ourTexture, TexCoord);
		}
	` + "\x00"

	// Basic vertex shader
	vertexShader2DSource = `
		#version 330 core
		layout (location = 0) in vec3 aPos;
		layout (location = 1) in vec2 aTexCoord;
		
		out vec2 TexCoord;

		uniform mat4 view;
		
		void main()
		{
			gl_Position = view * vec4(aPos, 1.0);
			TexCoord = aTexCoord;
		}
	` + "\x00"

	// Basic fragment shader
	fragmentShader2DSource = `
		#version 330 core
		out vec4 FragColor;
		
		in vec2 TexCoord;
		
		uniform sampler2D ourTexture;
		
		void main()
		{
			FragColor = texture(ourTexture, TexCoord);
		}
	` + "\x00"
)
