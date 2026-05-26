// Copyright 2019 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package exec

// TODO: we could use the Docker Go SDK for this, but it seems like a big dependency to pull in just
// to run some "docker exec" commands.

// RunDockerExecCommand runs the provided command on the specified host using "docker exec". Returns
// the exit code of the command, along with the contents of stdout and stderr as strings. Note that
// if the command returns a non-zero error code, this function does not report it as an error.
func RunDockerExecCommand(container, cmd, workdir string, envs map[string]string, stdin string) (
	code int, stdout string, stderr string, err error,
) {
	_ = "STUB: not implemented"
	return 0,

		// Set environment variables.
		"", "", nil
}

// Interactive mode to get input from stdin.

// Just split in to "/bin/sh" "-c" and "actual_cmd"
// This is useful for passing piped commands in to os/exec interface.

// command is successful

// RunDockerPsFilterCommand runs the provided command on the specified host using "docker ps filter". Returns
// the exit code of the command, along with the contents of stdout and stderr as strings.
func RunDockerPsFilterCommand(filter string) (
	code int, stdout string, stderr string, err error,
) {
	_ = "STUB: not implemented"
	return 0, "", "", nil
}

// command is successful
