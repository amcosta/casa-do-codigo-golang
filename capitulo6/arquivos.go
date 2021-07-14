package capitulo6

import (
	"fmt"
	"os"
)

func CriarArquivos() {
	tmp := os.TempDir()
	fmt.Println("Diretório temporario", tmp)

	criarArquivos(tmp)
	criarArquivos(tmp, "teste1")
	criarArquivos(tmp, "teste2", "teste3", "teste4")
}

func criarArquivos(dirBase string, arquivos ...string) {
	for _, nome := range arquivos {
		caminhoArquivo := fmt.Sprintf("%s/%s.%s", dirBase, nome, "txt")

		arq, err := os.Create(caminhoArquivo)
		if err != nil {
			fmt.Printf("Erro ao criar o arquivo %s: %v\n", caminhoArquivo, err)
			os.Exit(1)
		}

		defer arq.Close()
		fmt.Printf("Arquivo %s criado.\n", caminhoArquivo)
	}
}
