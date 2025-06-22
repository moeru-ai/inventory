package generator

import (
	"bytes"
	_ "embed"
	"html/template"
	"os"
	"path/filepath"

	"github.com/moeru-ai/inventory/cmd/tools/jem-generator/types"
	"github.com/nekomeowww/xo/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	//go:embed templates/providers.tmpl
	templateStrProviders string
	//go:embed templates/models.tmpl
	templateStrModels string
)

var (
	templateProviders = template.Must(template.New("providers").Parse(templateStrProviders))
	templateModels    = template.Must(template.New("models").Parse(templateStrModels))
)

var generatorLogger, _ = logger.NewLogger(
	logger.WithLevel(zapcore.DebugLevel),
	logger.WithCallFrameSkip(1),
	logger.WithNamespace("jem-generator"),
)

const fileMode = 0755

type Generator struct {
	providers []types.Provider
	models    []types.Model
	outputDir string
}

func New(providers []types.Provider, models []types.Model, outputDir string) (*Generator, error) {
	wd, err := os.Getwd()
	if err != nil {
		generatorLogger.Error("Error getting working directory.", zap.Error(err))
		return nil, err
	}

	var outputDirAbs string
	if filepath.IsAbs(outputDir) {
		outputDirAbs = outputDir
	} else {
		outputDirAbs = filepath.Join(wd, outputDir)
	}

	err = os.MkdirAll(outputDirAbs, fileMode)
	if err != nil {
		generatorLogger.Error("Error creating providers directory.", zap.Error(err))
		return nil, err
	}

	return &Generator{
		providers: providers,
		models:    models,
		outputDir: outputDirAbs,
	}, nil
}

func (g *Generator) generateProviders() ([]byte, error) {
	fileContent := bytes.NewBuffer(nil)
	err := templateProviders.Execute(fileContent, g.providers)

	if err != nil {
		generatorLogger.Error("Error executing template.", zap.Error(err))

		return nil, err
	}

	return fileContent.Bytes(), nil
}

func (g *Generator) generateModels() ([]byte, error) {
	fileContent := bytes.NewBuffer(nil)
	err := templateModels.Execute(fileContent, g.models)

	if err != nil {
		generatorLogger.Error("Error executing template.", zap.Error(err))

		return nil, err
	}

	return fileContent.Bytes(), nil
}

func (g *Generator) Generate() error {
	providersFilePath := filepath.Join(g.outputDir, "providers.ts")
	providersFileContent, err := g.generateProviders()

	if err != nil {
		return err
	}

	err = os.WriteFile(providersFilePath, providersFileContent, fileMode)
	if err != nil {
		return err
	}

	generatorLogger.With(
		zap.String("file_path", providersFilePath),
	).Info("File written.")

	modelsFilePath := filepath.Join(g.outputDir, "models.ts")
	modelsFileContent, err := g.generateModels()

	if err != nil {
		return err
	}

	err = os.WriteFile(modelsFilePath, modelsFileContent, fileMode)
	if err != nil {
		return err
	}

	generatorLogger.With(
		zap.String("file_path", modelsFilePath),
	).Info("File written.")

	return nil
}
