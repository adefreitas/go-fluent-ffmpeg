package fluentffmpeg

// GetInputPath returns the input file path
func (a *Args) GetInputPath() []string {
	var result []string
	for i := 0; i < len(a.inputs); i++ {
		if a.inputs[i].inputPath != "" {
			if a.inputs[i].nativeFramerateInput {
				result = append(result, "-re", "-i", a.inputs[i].inputPath)
			}
			result = append(result, "-i", a.inputs[i].inputPath)
		}
	}
	return result

}

// GetPipeInput returns whether or not ffmpeg is set to receive piped input
// func (a *Args) GetPipeInput() []string {
// 	if a.input.pipeInput != false {
// 		return []string{"-i", "pipe:0"}
// 	}

// 	return nil
// }

// // GetFromFormat returns the input format
// func (a *Args) GetFromFormat() []string {
// 	if a.input.fromFormat != "" {
// 		return []string{"-f", a.input.fromFormat}
// 	}

// 	return nil
// }

// // GetInputOptions returns additional input options
// func (a *Args) GetInputOptions() []string {
// 	return a.input.inputOptions
// }
