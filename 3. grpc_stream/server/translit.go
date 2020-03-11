package main

import (
	"coursera/microservices/grpc_stream/translit"
	"fmt"
	"io"

	tr "github.com/gen1us2k/go-translit"
)

type TrServer struct {
}

//Реализация сервиса, который был задан в *.proto файле
func (srv *TrServer) EnRu(inStream translit.Transliteration_EnRuServer) error {
	//запускается в одной горутине
	for {
		//получение данных с канала
		inWord, err := inStream.Recv()

		//проверка, закрылся ли стрим
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		//создание структуры для отправки на клиент
		out := &translit.Word{
			Word: tr.Translit(inWord.Word),
		}
		fmt.Println(inWord.Word, "->", out.Word)

		//отправка данных
		inStream.Send(out)
	}
	return nil
}

func NewTr() *TrServer {
	return &TrServer{}
}
