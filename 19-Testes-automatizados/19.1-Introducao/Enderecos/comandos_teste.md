# go test - opções e flags

- 1. ### go test
  - Modo normal, mostra se os testes passaram ou não
  ```Bash
  ❯ go test
  PASS
  ok      go_aprenda_do_zero/19-Testes-automatizados/Enderecos    0.002s

- 2. ### go test -v
  - Modo verbose, mostra as execuções (se tiver funções com t.Parallel() ele irá demonstrar as pausas e continuações também)
  ```Bash
  ❯ go test -v
  === RUN   TestTipoDeEnderecoV1
  --- PASS: TestTipoDeEnderecoV1 (0.00s)
  === RUN   TestTipoDeEndereco
  === PAUSE TestTipoDeEndereco
  === RUN   TestQualquer
  === PAUSE TestQualquer
  === CONT  TestTipoDeEndereco
  === CONT  TestQualquer
  --- PASS: TestTipoDeEndereco (0.00s)
  --- PASS: TestQualquer (0.00s)
  PASS
  ok      go_aprenda_do_zero/19-Testes-automatizados/Enderecos    0.002s

- 3. ### go test --cover
  - Mostra a cobertura de código
  ```Bash
  ❯ go test --cover
  PASS  
  coverage: 90.0% of statements
  ok      go_aprenda_do_zero/19-Testes-automatizados/Enderecos    0.002s

- 4. ### go test --coverprofile \<arquivo>.txt
  - Gera um arquivo .txt com a cobertura mais detalhada
  ```Bash
  ❯ go test --coverprofile cobertura.txt
  PASS
  coverage: 90.0% of statements
  ok      go_aprenda_do_zero/19-Testes-automatizados/Enderecos    0.003s

  ❯ cat cobertura.txt
  mode: set
  go_aprenda_do_zero/19-Testes-automatizados/Enderecos/enderecos.go:10.45,15.52 3 1
  go_aprenda_do_zero/19-Testes-automatizados/Enderecos/enderecos.go:15.52,17.3 1 1
  go_aprenda_do_zero/19-Testes-automatizados/Enderecos/enderecos.go:19.2,19.25 1 1
  go_aprenda_do_zero/19-Testes-automatizados/Enderecos/enderecos.go:22.34,23.17 1 1
  go_aprenda_do_zero/19-Testes-automatizados/Enderecos/enderecos.go:23.17,25.3 1 0
  go_aprenda_do_zero/19-Testes-automatizados/Enderecos/enderecos.go:26.2,28.22 3 1

- 5. ### go tool cover --func=\<arquivo>.txt
  - Demonstra no terminal uma análise do arquivo gerado pelo --coverprofile
  ```Bash
  ❯ go tool cover --func=cobertura.txt
  go_aprenda_do_zero/19-Testes-automatizados/Enderecos/enderecos.go:10:   TipoDeEndereco  100.0%
  go_aprenda_do_zero/19-Testes-automatizados/Enderecos/enderecos.go:22:   capitalize      80.0%
  total:   (statements)    90.0%

- 6. ### go tool cover --html=\<arquivo>.txt
  - Gera um HTML detalhando o que não está coberto pelos testes
  ```Bash
  ❯ go tool cover --html=cobertura.txt
