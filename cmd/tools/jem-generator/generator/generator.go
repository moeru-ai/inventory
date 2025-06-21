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
	//go:embed templates/index.tmpl
	templateStrIndex string
)

var (
	templateProviders = template.Must(template.New("providers").Parse(templateStrProviders))
	templateIndex     = template.Must(template.New("index").Parse(templateStrIndex))
)

var providerLogger, _ = logger.NewLogger(
	logger.WithLevel(zapcore.DebugLevel),
	logger.WithCallFrameSkip(1),
	logger.WithNamespace("jem-generator"),
)

const fileMode = 0755

type Generator struct {
	providers []types.Provider
	outputDir string
}

func New(providers []types.Provider, outputDir string) (*Generator, error) {
	wd, err := os.Getwd()
	if err != nil {
		providerLogger.Error("Error getting working directory.", zap.Error(err))
		return nil, err
	}

	var providersDir string
	if filepath.IsAbs(outputDir) {
		providersDir = filepath.Join(outputDir, "providers")
	} else {
		providersDir = filepath.Join(wd, outputDir, "providers")
	}

	err = os.MkdirAll(providersDir, fileMode)
	if err != nil {
		providerLogger.Error("Error creating providers directory.", zap.Error(err))
		return nil, err
	}

	return &Generator{
		providers: providers,
		outputDir: providersDir,
	}, nil
}

func (g *Generator) generateProvider(provider types.Provider) ([]byte, error) {
	fileContent := bytes.NewBuffer(nil)
	err := templateProviders.Execute(fileContent, provider)

	if err != nil {
		providerLogger.With(
			zap.String("provider", provider.Name),
		).Error("Error executing template.", zap.Error(err))

		return nil, err
	}

	return fileContent.Bytes(), nil
}

func (g *Generator) generateProviders() error {
	for _, provider := range g.providers {
		filePath := filepath.Join(g.outputDir, provider.Name+".ts")
		fileContent, err := g.generateProvider(provider)

		if err != nil {
			return err
		}

		err = os.WriteFile(filePath, fileContent, fileMode)
		if err != nil {
			return err
		}

		providerLogger.With(
			zap.String("file_path", filePath),
		).Info("File written.")
	}

	return nil
}

func (g *Generator) generateIndex() ([]byte, error) {
	fileContent := bytes.NewBuffer(nil)
	err := templateIndex.Execute(fileContent, g.providers)

	if err != nil {
		providerLogger.Error("Error executing template.", zap.Error(err))
		return nil, err
	}

	return fileContent.Bytes(), nil
}

func (g *Generator) Generate() error {
	err := g.generateProviders()
	if err != nil {
		return err
	}

	indexFilePath := filepath.Join(g.outputDir, "index.ts")
	indexFileContent, err := g.generateIndex()

	if err != nil {
		return err
	}

	err = os.WriteFile(indexFilePath, indexFileContent, fileMode)
	if err != nil {
		return err
	}

	providerLogger.With(
		zap.String("file_path", indexFilePath),
	).Info("File written.")

	return nil
}
