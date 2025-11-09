import './App.css'
import { Routes, Route } from 'react-router-dom';
import CadastroUsuario from '../src/components/CadastroUsuario.jsx'
import CadastroProduto from '../src/components/CadastroProduto.jsx';
import CriarPedido from '../src/components/CriarPedido.jsx';

function App() {
    return (
        <Routes>
            <Route path="/" element={<CadastroUsuario />} />
            <Route path="/produto" element={<CadastroProduto />} />
            <Route path="/pedido" element={<CriarPedido />} />
        </Routes>
    );
}

export default App;