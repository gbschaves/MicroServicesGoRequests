import React, { useState } from 'react';

const CadastroProduto = () => {
    // Estado para o Nome e Preço
    const [nomeProduto, setNomeProduto] = useState('');
    const [preco, setPreco] = useState('');

    const handleSubmit = (event) => {
        // Previne o recarregamento da página
        event.preventDefault();

        // Dados a serem enviados:
        console.log('Dados do Produto a serem enviados:', {
            nome: nomeProduto,
            preco: preco
        });

        // ⚠️ Lógica de conversão de preço (ex: "R$ 10,50" para 10.50) deve vir aqui antes da API.
        // api.post('/produtos', { nome: nomeProduto, preco: precoNumerico });
    };

    return (
        <div className="relative flex h-screen w-full flex-col">
            <header className="flex shrink-0 items-center bg-background-light dark:bg-background-dark p-4 pb-2">
                <button className="text-gray-900 dark:text-white flex size-12 items-center justify-center">
                    <span className="material-symbols-outlined text-2xl">arrow_back</span>
                </button>
                <h1 className="text-gray-900 dark:text-white text-lg font-bold leading-tight tracking-[-0.015em] flex-1 text-center pr-12">Cadastrar Produto</h1>
            </header>

            {/* O Formulário agora engloba tanto os campos quanto o botão de submissão */}
            <form
                onSubmit={handleSubmit}
                className="flex-1 overflow-y-auto flex flex-col justify-between"
            >
                <div className="flex flex-col gap-6 px-4 py-3">
                    <label className="flex flex-col">
                        <p className="text-gray-800 dark:text-white text-base font-medium leading-normal pb-2">Nome do Produto</p>
                        <input
                            className="form-input flex w-full min-w-0 flex-1 resize-none overflow-hidden rounded-lg text-gray-900 dark:text-white focus:outline-0 focus:ring-2 focus:ring-primary border-gray-300 dark:border-gray-700 bg-white dark:bg-[#233648] h-14 placeholder:text-gray-400 dark:placeholder:text-[#92adc9] p-4 text-base font-normal leading-normal"
                            placeholder="Digite o nome do produto"
                            value={nomeProduto}
                            onChange={(e) => setNomeProduto(e.target.value)}
                        />
                    </label>
                    <label className="flex flex-col">
                        <p className="text-gray-800 dark:text-white text-base font-medium leading-normal pb-2">Preço</p>
                        <input
                            className="form-input flex w-full min-w-0 flex-1 resize-none overflow-hidden rounded-lg text-gray-900 dark:text-white focus:outline-0 focus:ring-2 focus:ring-primary border-gray-300 dark:border-gray-700 bg-white dark:bg-[#233648] h-14 placeholder:text-gray-400 dark:placeholder:text-[#92adc9] p-4 text-base font-normal leading-normal"
                            inputMode="decimal"
                            placeholder="R$ 0,00"
                            type="text"
                            value={preco}
                            onChange={(e) => setPreco(e.target.value)}
                        />
                    </label>
                </div>

                {/* Rodapé (agora dentro do <form>) */}
                <footer className="p-4 bg-background-light dark:bg-background-dark">
                    <button
                        type="submit" // Aciona o onSubmit do formulário pai
                        className="flex w-full cursor-pointer items-center justify-center overflow-hidden rounded-lg h-12 px-5 bg-primary text-white text-base font-bold leading-normal tracking-[0.015em] hover:bg-primary/90 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-background-dark focus:ring-primary"
                    >
                        <span className="truncate">Adicionar Produto</span>
                    </button>
                </footer>
            </form>
        </div>
    );
};

export default CadastroProduto;