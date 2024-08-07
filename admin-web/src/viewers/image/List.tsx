import { useRequest } from "ahooks";
import { getImageList, uploadImageToGithubRepo } from "../../actions";

function ImageList() {
    const { data } = useRequest(() => getImageList(), { refreshDeps: [] });
    console.log(data);
    const uploadFile = async () => {
        const input = document.createElement("input");
        input.type = "file";
        input.accept = "image/*";
        input.style.display = "none";
        input.onchange = async (e) => {
            const file = (e.target as HTMLInputElement).files?.[0];
            if (file) {
                const formData = new FormData();
                formData.append("file", file);
                uploadImageToGithubRepo(formData);
            }
        };
        setTimeout(() => {
            document.body.appendChild(input);
            input.click();
        }, 0);
    };

    return (
        <>
            <button onClick={uploadFile}>upload file to github repo</button>
            <table>
                <thead>
                    <tr>
                        <th>id</th>
                        <th>thumb</th>
                        <th>width</th>
                        <th>height</th>
                        <th>size</th>
                        <th>create_at</th>
                        <th>update_at</th>
                    </tr>
                </thead>
                <tbody>
                    {data?.map((item) => {
                        return (
                            <tr key={item.id}>
                                <td>{item.id}</td>
                                <td>
                                    <img src={item.url} width={200} />
                                </td>
                                <td>{item.width}</td>
                                <td>{item.height}</td>
                                <td>{item.size}</td>
                                <td>{item.create_at}</td>
                                <td>{item.updated_at}</td>
                            </tr>
                        );
                    })}
                </tbody>
            </table>
        </>
    );
}

export default ImageList;
