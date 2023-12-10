"use server";

const send_email_api_url = "http://localhost:8080/send"
//const send_email_api_url = "https://dearself.onrender.com/send"

export async function create(formData: FormData): Promise<any> {
        "use server";
        const response = await fetch(send_email_api_url, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                
            },
            body: JSON.stringify({
                name: formData.get("name"),
                to: formData.get("to"),
                subject: formData.get("subject"),
                text: formData.get("text"),
                date: formData.get("date"),
            }),
        });

        return response.json();
}
