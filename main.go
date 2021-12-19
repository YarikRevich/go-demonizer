package godemonizer

import (
	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(os.Stderr)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}
