package generator

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/go-openapi/analysis"
	"github.com/go-openapi/loads"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSpec_Issue1429(t *testing.T) {
	log.SetOutput(io.Discard)
	defer log.SetOutput(os.Stdout)

	// acknowledge fix in go-openapi/spec
	specPath := filepath.Join("..", "fixtures", "bugs", "1429", "swagger-1429.yaml")
	_, err := loads.Spec(specPath)
	assert.NoError(t, err)

	opts := testGenOpts()
	opts.Spec = specPath
	_, err = opts.validateAndFlattenSpec()
	assert.NoError(t, err)

	// more aggressive fixture on $refs, with validation errors, but flatten ok
	specPath = filepath.Join("..", "fixtures", "bugs", "1429", "swagger.yaml")
	specDoc, err := loads.Spec(specPath)
	assert.NoError(t, err)

	opts.Spec = specPath
	opts.FlattenOpts.BasePath = specDoc.SpecFilePath()
	opts.FlattenOpts.Spec = analysis.New(specDoc.Spec())
	opts.FlattenOpts.Minimal = true
	err = analysis.Flatten(*opts.FlattenOpts)
	assert.NoError(t, err)

	specDoc, _ = loads.Spec(specPath) // needs reload
	opts.FlattenOpts.Spec = analysis.New(specDoc.Spec())
	opts.FlattenOpts.Minimal = false
	err = analysis.Flatten(*opts.FlattenOpts)
	assert.NoError(t, err)
}

func TestSpec_Issue2527(t *testing.T) {
	log.SetOutput(io.Discard)
	defer log.SetOutput(os.Stdout)

	t.Run("spec should be detected as invalid", func(t *testing.T) {
		specPath := filepath.Join("..", "fixtures", "bugs", "2527", "swagger.yml")
		_, err := loads.Spec(specPath)
		require.NoError(t, err)

		opts := testGenOpts()
		opts.Spec = specPath
		opts.ValidateSpec = true // test options skip validation by default
		_, err = opts.validateAndFlattenSpec()
		require.Error(t, err)
	})

	t.Run("fixed spec should be detected as valid", func(t *testing.T) {
		specPath := filepath.Join("..", "fixtures", "bugs", "2527", "swagger-fixed.yml")
		_, err := loads.Spec(specPath)
		require.NoError(t, err)

		opts := testGenOpts()
		opts.Spec = specPath
		opts.ValidateSpec = true
		_, err = opts.validateAndFlattenSpec()
		require.NoError(t, err)
	})
}

func TestSpec_FindSwaggerSpec(t *testing.T) {
	keepErr := func(_ string, err error) error { return err }
	assert.Error(t, keepErr(findSwaggerSpec("")))
	assert.Error(t, keepErr(findSwaggerSpec("nowhere")))
	assert.Error(t, keepErr(findSwaggerSpec(filepath.Join("..", "fixtures"))))
	assert.NoError(t, keepErr(findSwaggerSpec(filepath.Join("..", "fixtures", "codegen", "shipyard.yml"))))
}

func TestSpec_Issue1621(t *testing.T) {
	log.SetOutput(io.Discard)
	defer log.SetOutput(os.Stdout)

	// acknowledge fix in go-openapi/spec
	specPath := filepath.Join("..", "fixtures", "bugs", "1621", "fixture-1621.yaml")
	_, err := loads.Spec(specPath)
	require.NoError(t, err)

	opts := testGenOpts()
	opts.Spec = specPath
	opts.ValidateSpec = true
	_, err = opts.validateAndFlattenSpec()
	assert.NoError(t, err)
}

func TestShared_Issue1614(t *testing.T) {
	log.SetOutput(io.Discard)
	defer log.SetOutput(os.Stdout)

	// acknowledge fix in go-openapi/spec
	specPath := filepath.Join("..", "fixtures", "bugs", "1614", "gitea.json")
	_, err := loads.Spec(specPath)
	require.NoError(t, err)

	opts := testGenOpts()
	opts.Spec = specPath
	opts.ValidateSpec = true
	_, err = opts.validateAndFlattenSpec()
	assert.NoError(t, err)
}

func Test_analyzeSpec_Issue2216(t *testing.T) {
	log.SetOutput(io.Discard)
	defer log.SetOutput(os.Stdout)

	t.Run("single-swagger-file", func(t *testing.T) {
		specPath := filepath.Join("..", "fixtures", "bugs", "2216", "swagger-single.yml")

		opts := testGenOpts()
		opts.Spec = specPath
		opts.ValidateSpec = true
		opts.PropertiesSpecOrder = true
		_, _, err := opts.analyzeSpec()
		assert.NoError(t, err)
	})

	t.Run("splitted-swagger-file", func(t *testing.T) {
		specPath := filepath.Join("..", "fixtures", "bugs", "2216", "swagger.yml")

		opts := testGenOpts()
		opts.Spec = specPath
		opts.ValidateSpec = true
		opts.PropertiesSpecOrder = true
		_, _, err := opts.analyzeSpec()
		assert.NoError(t, err)
	})
}
