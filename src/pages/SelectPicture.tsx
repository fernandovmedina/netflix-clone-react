const SelectPicture = () => {
    document.addEventListener("DOMContentLoaded", async () => {
        const URL: string = "http://127.0.0.1:8010/microservice/titles";
        const titlesList: any = document.getElementById("titles-list");

        try {
            const response = await fetch(URL, { method: "GET" });
            const result = await response.json();
            const data = result.body;

            data.forEach((item: any) => {
                const listItem = document.createElement("li");
                listItem.innerHTML = `<strong>Name:</strong> ${item.Name} - <strong>Length:</strong> ${item.Len}`;
                titlesList.appendChild(listItem);
            });
        } catch (err) {
            console.error("Failed to fetch titles:", err);
        }
    });

    return (
        <main>
            <section className="bg-gradient-to-b from-black to-gray-950">
                <section
                    id="titles-section"
                    className="bg-gradient-to-b from-black to-gray-950"
                >
                    <ul id="titles-list"></ul>
                </section>
            </section>
        </main>
    )
}

export default SelectPicture;
