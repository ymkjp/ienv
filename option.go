package ienv

import (
  "flag"
  //"fmt"
  "os"
  "io/ioutil"
  uuid "github.com/google/uuid"
  log "github.com/sirupsen/logrus"
)

type Option struct {
  Url string
  Dir string
  DeployTo string
  Debug bool
  Help bool
  Uuid *uuid.UUID
}

func (opt *Option) init() {
  homeDir := os.Getenv("HOME")
  opt.Uuid = new(uuid.UUID)
  opt.setupFlag()
  flag.Parse()
  if opt.Url == "" {
    flag.PrintDefaults()
    log.Fatal("Specify required option: --url")
  }
  if opt.Dir == "" {
    opt.Dir = opt.targetDir(homeDir, opt.Uuid.String())
  }
  if opt.DeployTo == "" && homeDir == "" {
    log.Fatal("Any of environment variable $HOME or --deploy-to needs to be specified")
  } else if opt.DeployTo == "" {
    opt.DeployTo = opt.targetDir(homeDir, opt.Uuid.String())
  }
}

func (opt *Option) setupFlag() {
  flag.StringVar(
    &opt.Url,
    "url",
    "",
    "[Required] URL to dotfiles repository such as https://github.com/ymkjp/dotfiles")
  flag.StringVar(
    &opt.Dir,
    "dir",
    "",
    "[Optional] Directory path used by git clone. " +
      "If nothing specified, these paths will be used by following order:\n\t" +
      "1. Environment variable $HOME\n\t" +
      "2. Temporary directory")
  flag.StringVar(
    &opt.DeployTo,
    "deploy-to",
    "",
    "[Optional] Directory path which symlinks of dotfiles is going to be deployed. " +
      "If nothing specified, environment variable $HOME will be used.")
  flag.BoolVar(
    &opt.Debug,
    "debug",
    false,
    "[Optional] Debug mode. Temporary directory will be used as --dir and --deploy-to.")
}

func (opt *Option) targetDir(dir string, prefix string) string {
  var err error
  if opt.Debug || dir == "" {
    dir, err = ioutil.TempDir("", prefix)
  }
  if err != nil {
    log.Fatal(err)
  }
  return dir
}
