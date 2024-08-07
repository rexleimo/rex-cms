import { Route, Routes } from "react-router-dom";
import ImageList from "./List";

function ImagePage() {
    return (
        <Routes>
            <Route path='/' element={<ImageList />} />
        </Routes>
    );
}

export default ImagePage;
