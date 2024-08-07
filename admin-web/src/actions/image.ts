import ajax from "./request";

interface Image {
    id: number;
    url: string;
    width: number;
    height: number;
    size: number;
    mime_type: string;
    create_at: string;
    updated_at: string;
}

export function getImageList(): Promise<Image[]> {
    return ajax({
        url: "/image",
        method: "GET",
    }).then((res) => res.data);
}

export function uploadImageToGithubRepo(formData: FormData): Promise<unknown> {
    return ajax({
        url: "/image/publish",
        method: "POST",
        data: formData,
        headers: {
            "Content-Type": "multipart/form-data",
        },
    });
}
