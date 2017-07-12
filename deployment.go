package ienv

import (
  "path/filepath"
  "os"
  log "github.com/sirupsen/logrus"
  "path"
)

type Deployment struct {
  SourcePath string
  DestinationPath string
  CopyPattern string
  IgnorePattern string
}

func (dep *Deployment) setTarget(srcPath string, dstPath string) *Deployment {
  dep.SourcePath = srcPath
  dep.DestinationPath = dstPath
  return dep
}

func (dep *Deployment) setPattern(ignorePattern string, copyPattern string) *Deployment {
  dep.IgnorePattern = ignorePattern
  dep.CopyPattern = copyPattern
  return dep
}

func (dep *Deployment) sync() {
  err := filepath.Walk(dep.SourcePath, dep.visit)
  if err != nil {
    log.WithFields(log.Fields {
      "SourcePath": dep.SourcePath,
    }).Error("Error occurred during walking")
    log.Fatal(err)
  }
}

func (dep *Deployment) visit(currentPath string, info os.FileInfo, err error) error {
  if err != nil {
    log.Fatal(err)
  }
  // @todo Adopt Patterns
  relPath, err := filepath.Rel(dep.SourcePath, currentPath)
  if err != nil {
    log.Fatal(err)
  }
  candidatePath := path.Join(dep.DestinationPath, relPath)
  err = dep.ensureSymlink(currentPath, candidatePath)
  if err != nil {
    log.Fatal(err)
  }
  return nil
}

func (dep *Deployment) ensureSymlink(src string, dst string) error {
  logger := log.WithFields(log.Fields {
    "src": src,
    "dst": dst,
  })
  if _, err := os.Stat(dst); os.IsNotExist(err) {
    err := os.Symlink(src, dst)
    if err != nil {
      logger.Error("Failed to ensure symlink")
      return err
    }
    logger.Info("Successfully created symlink")
  } else {
    logger.Info("Already deployed")
  }
  return nil
}
