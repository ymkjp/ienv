package ienv

func Run() {
  opt := new(Option)
  src := new(Source)
  opt.init()
  src.fetch(opt.Url, opt.Dir)
  src.cleanup(opt.Dir, !opt.Debug)
}
