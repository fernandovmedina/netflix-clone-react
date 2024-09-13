import { useState, useEffect } from "react";
import FooterSign from "../components/FooterSign";
import NavbarSign from "../components/NavbarSign";

interface Profile {
    id: number;
    name: string;
    type: string;
}

const CreateProfile = () => {
    const [profiles, setProfiles] = useState<Profile[]>([]);
    const email: string | null = new URLSearchParams(window.location.search).get("email");
    const password: string | null = new URLSearchParams(window.location.search).get("password");
    const plan: string | null = new URLSearchParams(window.location.search).get("plan");
    const cardNumber: string | null = new URLSearchParams(window.location.search).get("cardNumber");
    const dueDate: string | null = new URLSearchParams(window.location.search).get("dueDate");
    const cvv: string | null = new URLSearchParams(window.location.search).get("cvv");
    const cardName: string | null = new URLSearchParams(window.location.search).get("name");

    useEffect(() => {
        const spawns = document.getElementById("spawns");
        const add = document.getElementById("add");
        let counter = 0;

        const handleAddClick = () => {
            counter++;
            const newProfile: Profile = { id: counter, name: "", type: "adults" };
            setProfiles((prevProfiles) => [...prevProfiles, newProfile]);
        };

        add?.addEventListener("click", handleAddClick);

        return () => {
            add?.removeEventListener("click", handleAddClick);
        };
    }, []);


    const handleProfileChange = (id: number, field: "name" | "type", value: string) => {
        setProfiles((prevProfiles) =>
            prevProfiles.map((profile) =>
                profile.id === id ? { ...profile, [field]: value } : profile
            )
        );
    };

    const handlePost = () => {
        const users = profiles.reduce((acc, profile, index) => {
            acc[index] = {
                name: profile.name,
                type: profile.type
            };
            return acc;
        }, {});
    
        const json = {
            body: {
                email: email,
                password: password,
                plan: plan,
                users: users,
                payload: {
                    type: "card",
                    card_number: cardNumber,
                    due_date: dueDate,
                    cvv: cvv,
                    card_name: cardName
                }
            }
        };
    
        console.log(JSON.stringify(json))

        fetch("http://127.0.0.1:8030/microservice/register", {
            method: "POST",
            headers: {"Content-Type": "application/json"},
            body: JSON.stringify(json),
        }).then((res) => res.json()).then((res) => console.log(res)).catch((err) => console.error(err))
    };
    

    return (
        <main>
            <NavbarSign />

            <section className="w-full flex items-center justify-center">
                <div className="w-[35%] py-16">
                    <div className="flex items-center justify-between">
                        <h1 className="text-4xl font-bold mb-4">Create profiles</h1>
                        <div>
                            <img
                                id="add"
                                className="w-10 hover:w-12"
                                src="/add_icon_red.png"
                                alt="add_icon_red_image"
                            />
                        </div>
                    </div>

                    {profiles.map((profile) => (
                        <div key={profile.id} className="flex my-2">
                            <img
                                className="w-16"
                                src="/netflix_profile_image_yellow.jpg"
                                alt="netflix_profile_image_yellow_image"
                            />
                            <input
                                className="mx-4 border-2 border-gray-500 rounded px-3 w-full"
                                type="text"
                                placeholder="Profile name"
                                value={profile.name}
                                onChange={(e) => handleProfileChange(profile.id, "name", e.target.value)}
                            />
                            <select
                                name="ages"
                                id={`ages-${profile.id}`}
                                className="bg-white font-bold text-lg"
                                value={profile.type}
                                onChange={(e) => handleProfileChange(profile.id, "type", e.target.value)}
                            >
                                <option value="adults">Adults</option>
                                <option value="kids">Kids</option>
                            </select>
                        </div>
                    ))}

                    <button
                        onClick={handlePost}
                        className="bg-red-600 w-full my-4 text-white py-3 rounded hover:bg-red-700"
                    >
                        SAVE
                    </button>
                </div>
            </section>

            <FooterSign />
        </main>
    );
};

export default CreateProfile;
