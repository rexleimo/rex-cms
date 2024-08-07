import Image from "./viewers/image";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import { Login } from "./viewers";
function App() {
    return (
        <BrowserRouter>
            <Routes>
                <Route path='/login' element={<Login />} />
                <Route path='/image/*' element={<Image />} />
            </Routes>
        </BrowserRouter>
    );
}

export default App;
