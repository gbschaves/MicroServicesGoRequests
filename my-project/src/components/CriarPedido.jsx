import React, { useState } from 'react';

const CriarPedido = () => {
    // 1. Estado para os campos do formulário
    const [idUsuario, setIdUsuario] = useState('');
    const [idProduto, setIdProduto] = useState('');
    const [quantidade, setQuantidade] = useState(1);

    // 2. Função de submissão
    const handleSubmit = (event) => {
        event.preventDefault();

        // Dados do pedido a serem enviados:
        console.log('Dados do Pedido a serem enviados:', {
            idUsuario: idUsuario,
            idProduto: idProduto,
            quantidade: Number(quantidade) // Converte a string de quantidade para número
        });

        // ⚠️ Aqui você faria a chamada para a API (ex: api.post('/pedidos', {...}))

        // Opcional: Limpar o formulário ou redirecionar o usuário
    };

    return (
        <div className="relative flex h-auto min-h-screen w-full flex-col dark group/design-root overflow-x-hidden" style={{ '--select-button-svg': "url('data:image/svg+xml,%3csvg xmlns=%27http://www.w3.org/2000/svg%27 width=%2724px%27 height=%2724px%27 fill=%27rgb(146,173,201)%27 viewBox=%270 0 256 256%27%3e%3cpath d=%27M181.66,170.34a8,8,0,0,1,0,11.32l-48,48a8,8,0,0,1-11.32,0l-48-48a8,8,0,0,1,11.32-11.32L128,212.69l42.34-42.35A8,8,0,0,1,181.66,170.34Zm-96-84.68L128,43.31l42.34,42.35a8,8,0,0,0,11.32-11.32l-48-48a8,8,0,0,0-11.32,0l-48-48A8,8,0,0,0,85.66,85.66Z%27%3e%3c/path%3e%3c/svg%3e')" }}>
            <div className="flex flex-col flex-1 h-full">
                <header className="flex items-center p-4 pb-2 justify-between sticky top-0 bg-background-light dark:bg-background-dark z-10">
                    <div className="flex size-12 shrink-0 items-center text-gray-800 dark:text-white">
                        <span className="material-symbols-outlined !text-2xl">arrow_back</span>
                    </div>
                    <h1 className="text-gray-900 dark:text-white text-lg font-bold leading-tight tracking-[-0.015em] flex-1">Criar Novo Pedido</h1>
                    <div className="flex size-12 shrink-0 items-center"></div>
                </header>

                {/* 3. O formulário agora envolve o conteúdo e o rodapé */}
                <form onSubmit={handleSubmit} className="flex flex-col flex-1">

                    <main className="flex-1 px-4 py-3">
                        <div className="flex flex-col gap-6">

                            {/* Campo ID do Usuário */}
                            <label className="flex flex-col">
                                <p className="text-gray-800 dark:text-white text-base font-medium leading-normal pb-2">ID do Usuário</p>
                                <div className="relative">
                                    <input
                                        className="flex w-full min-w-0 flex-1 resize-none overflow-hidden rounded-lg text-gray-800 dark:text-white focus:outline-0 focus:ring-2 focus:ring-primary/50 border-gray-300 dark:border-gray-700 bg-white dark:bg-gray-800/50 focus:border-primary h-14 placeholder:text-gray-400 dark:placeholder:text-gray-500 p-[15px] text-base font-normal leading-normal pr-4"
                                        placeholder="Digite o ID do usuário"
                                        type="text"
                                        value={idUsuario}
                                        onChange={(e) => setIdUsuario(e.target.value)}
                                    />
                                </div>
                            </label>

                            {/* Campo ID do Produto */}
                            <label className="flex flex-col">
                                <p className="text-gray-800 dark:text-white text-base font-medium leading-normal pb-2">ID do Produto</p>
                                <div className="relative">
                                    <input
                                        className="flex w-full min-w-0 flex-1 resize-none overflow-hidden rounded-lg text-gray-800 dark:text-white focus:outline-0 focus:ring-2 focus:ring-primary/50 border-gray-300 dark:border-gray-700 bg-white dark:bg-gray-800/50 focus:border-primary h-14 placeholder:text-gray-400 dark:placeholder:text-gray-500 p-[15px] text-base font-normal leading-normal pr-4"
                                        placeholder="Digite o ID do produto"
                                        type="text"
                                        value={idProduto}
                                        onChange={(e) => setIdProduto(e.target.value)}
                                    />
                                </div>
                            </label>

                            {/* Campo Quantidade */}
                            <label className="flex flex-col">
                                <p className="text-gray-800 dark:text-white text-base font-medium leading-normal pb-2">Quantidade</p>
                                <div className="relative">
                                    <input
                                        className="flex w-full min-w-0 flex-1 resize-none overflow-hidden rounded-lg text-gray-800 dark:text-white focus:outline-0 focus:ring-2 focus:ring-primary/50 border-gray-300 dark:border-gray-700 bg-white dark:bg-gray-800/50 focus:border-primary h-14 placeholder:text-gray-400 dark:placeholder:text-gray-500 p-[15px] text-base font-normal leading-normal pr-4 [appearance:textfield] [&amp;::-webkit-inner-spin-button]:appearance-none [&amp;::-webkit-outer-spin-button]:appearance-none"
                                        min="1"
                                        placeholder="Digite a quantidade"
                                        type="number"
                                        value={quantidade}
                                        onChange={(e) => setQuantidade(e.target.value)}
                                    />
                                </div>
                            </label>
                        </div>
                    </main>

                    {/* Rodapé e botão de submissão (agora dentro do form) */}
                    <footer className="sticky bottom-0 p-4 bg-background-light dark:bg-background-dark">
                        <button
                            type="submit" // Garante que o handleSubmit será chamado
                            className="flex min-w-[84px] w-full max-w-full cursor-pointer items-center justify-center overflow-hidden rounded-lg h-12 px-5 bg-primary text-white text-base font-bold leading-normal tracking-[0.015em] hover:bg-primary/90 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-background-dark focus:ring-primary transition-colors"
                        >
                            <span className="truncate">Criar Pedido</span>
                        </button>
                    </footer>
                </form>
            </div>
        </div>
    );
};

export default CriarPedido;