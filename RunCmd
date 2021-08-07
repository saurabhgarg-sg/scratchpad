/*
RunCmd provides a standardised way to run OS commands across packages.
Parameters:
    workDir of the directory where command should be run. Defaults to current directory.
    name of the binary to execute.
    arguments to pass as string.
It returns the output of the command from the OS.
*/
func RunCmd(workDir string, binary string, args string) (Output, error) {
    parts := strings.Split(args, " ")
    if workDir == "" {
        workDir = os.Getenv("PWD")
    }
    // Start constructing the command struct for exec.
    cmdString := fmt.Sprintf("Running command '%s' with arguments '%s' in directory %s\n",
        binary, args, workDir)
    allOutput := Output{Cmd: cmdString}
    cmd := exec.Command(binary, parts...)
    cmd.Dir = workDir

    // Start capture of the output and error streams.
    stdoutIn, err := cmd.StdoutPipe()
    if err != nil {
        return allOutput, errors.Wrap(err, "Error from StdoutPipe()")
    }
    stderrIn, err := cmd.StderrPipe()
    if err != nil {
        return allOutput, errors.Wrap(err, "Error from StderrPipe()")
    }

    // Display & capture for verification the characters in the output & error stream.
    var stdoutBuf, stderrBuf bytes.Buffer
    stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
    stderr := io.MultiWriter(os.Stderr, &stderrBuf)
    err = cmd.Start()
    if err != nil {
        return allOutput, errors.Wrap(err, "cmd.Start() failed")
    }

    // Wait for command execution to finish.
    var wg sync.WaitGroup
    wg.Add(1)
    var errStdout, errStderr error

    go func() {
        _, errStdout = io.Copy(stdout, stdoutIn)
        wg.Done()
    }()

    _, errStderr = io.Copy(stderr, stderrIn)
    wg.Wait()
    if errStdout != nil || errStderr != nil {
        return allOutput, errors.Wrap(err, "stdout/stderr capture failed")
    }

    err = cmd.Wait()
    if err != nil {
        return allOutput, errors.Wrap(err, "cmd.Wait() failed")
    }

    allOutput.Output = stdoutBuf.String()
    allOutput.Errors = stderrBuf.String()
    return allOutput, nil
}
