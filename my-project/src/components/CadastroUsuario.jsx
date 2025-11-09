import React, { useState } from 'react';

const CadastroUsuario = () => {
    // 1. Definição do estado para os campos do formulário
    const [nome, setNome] = useState('');
    const [email, setEmail] = useState('');

    // 2. Função de submissão do formulário
    const handleSubmit = (event) => {
        // Previne o comportamento padrão de recarregar a página
        event.preventDefault();

        console.log('Dados do usuário a serem enviados:', {
            nome: nome,
            email: email
        });

        // ⚠️ Aqui você faria a chamada para a API (ex: axios.post('/api/users', { nome, email }))

        // Opcional: Limpar o formulário após o envio
        // setNome('');
        // setEmail('');
    };

    return (
        <div className="relative flex h-auto min-h-screen w-full flex-col font-display dark group/design-root overflow-x-hidden">
            <div className="flex w-full grow items-center justify-center dark:bg-background-dark py-12">
                <div className="w-full max-w-md p-4">
                    <div className="flex justify-center pb-8">
                        <div className="flex h-16 w-16 items-center justify-center rounded-2xl bg-primary">
                            <span className="material-symbols-outlined text-white text-4xl">
                                hexagon
                            </span>
                        </div>
                    </div>
                    <h1 className="text-gray-800 dark:text-white tracking-tight text-[32px] font-bold leading-tight text-center pb-6 pt-0 px-4">Crie sua conta</h1>

                    {/* O formulário envolve os campos e chama handleSubmit na submissão */}
                    <form onSubmit={handleSubmit}>
                        <div className="flex flex-col gap-4 px-4">

                            {/* Campo Nome */}
                            <label className="flex flex-col w-full">
                                {/* Ajuste: Removido o 'dark:text-white' duplicado no <p> */}
                                <p className="text-gray-600 dark:text-white text-base font-medium leading-normal pb-2">Nome</p>
                                <input
                                    className="form-input flex w-full min-w-0 flex-1 resize-none overflow-hidden rounded-lg text-gray-800 dark:text-white focus:outline-0 focus:ring-2 focus:ring-primary/50 border border-gray-300 dark:border-[#324d67] bg-white dark:bg-[#192633] h-14 placeholder:text-gray-400 dark:placeholder:text-[#92adc9] p-[15px] text-base font-normal leading-normal transition-all"
                                    placeholder="Digite seu nome"
                                    value={nome}
                                    onChange={(e) => setNome(e.target.value)}
                                />
                            </label>

                            {/* Campo Email */}
                            <label className="flex flex-col w-full">
                                {/* Ajuste: Removido o 'dark:text-white' duplicado no <p> */}
                                <p className="text-gray-600 dark:text-white text-base font-medium leading-normal pb-2">Email</p>
                                <input
                                    className="form-input flex w-full min-w-0 flex-1 resize-none overflow-hidden rounded-lg text-gray-800 dark:text-white focus:outline-0 focus:ring-2 focus:ring-primary/50 border border-gray-300 dark:border-[#324d67] bg-white dark:bg-[#192633] h-14 placeholder:text-gray-400 dark:placeholder:text-[#92adc9] p-[15px] text-base font-normal leading-normal transition-all"
                                    placeholder="Digite seu email"
                                    type="email"
                                    value={email}
                                    onChange={(e) => setEmail(e.target.value)}
                                />
                            </label>
                        </div>

                        <div className="flex flex-col items-center gap-4 px-4 pt-8">
                            {/* O botão aciona a submissão do formulário */}
                            <button
                                type="submit"
                                className="flex h-14 w-full items-center justify-center rounded-lg bg-primary px-6 text-base font-bold text-white shadow-sm transition-all hover:bg-primary/90 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary"
                            >
                                Cadastrar
                            </button>
                        </div>
                    </form>

                </div>
            </div>
        </div>
    );
};

export default CadastroUsuario;