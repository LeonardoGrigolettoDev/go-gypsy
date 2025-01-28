package loader

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

func CreateFile(file multipart.FileHeader, path string) error {
	srcFile, err := file.Open()
	if err != nil {
		return fmt.Errorf("erro ao abrir o arquivo: %w", err)
	}
	defer srcFile.Close()

	// Criar o arquivo no caminho especificado
	dstFile, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("erro ao criar o arquivo: %w", err)
	}
	defer dstFile.Close()

	// Copiar o conteúdo do arquivo recebido para o arquivo no sistema
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		fmt.Errorf("erro ao copiar o conteúdo para o arquivo: %w", err)
		return err
	}

	fmt.Println("Arquivo criado com sucesso:", path)
	return nil
}

// func LoadCSV(file io.Reader) {
// 	file, err := os.Open(file)
// }
