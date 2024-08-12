import { useEffect } from "react";
import ProfileCardsManage from "./layouts/ProfileCardsManage";

const ManageProfiles = () => {
    useEffect(() => {
        document.getElementById("btn")!.addEventListener("click", () => {
            window.location.href = "/browse";
        });
    }, []);

    return (
        <main>
            <section
                className="bg-gradient-to-b from-black to-gray-950 text-white h-[100vh] flex items-center justify-center"
            >
                <div className="text-center">
                    <h1 className="font-bold text-5xl">Manage profiles:</h1>
                    <ProfileCardsManage />
                    <button
                        id="btn"
                        className="border-2 border-gray-400 hover:border-white hover:text-white px-3 py-2 text-gray-400"
                    >READY</button>
                </div>
            </section>
        </main>
    )
}

export default ManageProfiles;
