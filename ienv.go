package ienv


type ExitCode int

func Run() ExitCode {
  opt := new(Option)
  src := new(Source)
  dep := new(Deployment)
  opt.init()
  src.fetch(opt.Url, opt.Dir)
  dep.setTarget(opt.Dir, opt.DeployTo).
    setPattern(opt.IgnorePattern, opt.CopyPattern).
    sync()
  if opt.Debug {
    src.cleanup(opt.Dir)
    //src.cleanup(opt.DeployTo)
  }
  return 0
}
