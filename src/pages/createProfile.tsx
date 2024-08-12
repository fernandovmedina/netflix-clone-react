import React, { useEffect, useState } from "react";
import FooterSign from "../components/FooterSign";
import NavbarSign from "../components/NavbarSign";

const CreateProfile = () => {
    useEffect(() => {
        const spawns = document.getElementById("spawns");
        const add = document.getElementById("add");
        let counter = 0;

        const handleAddClick = () => {
            counter++;

            const profileDiv = `
                <div id="div-${counter}" class="flex my-2">
                    <img class="w-16" src="/netflix_profile_image_yellow.jpg" alt="netflix_profile_image_yellow_image">
                    <input class="mx-4 border-2 border-gray-500 rounded px-3 w-full" type="text" placeholder="Profile name">
                    <select name="ages" id="ages-${counter}" class="bg-white font-bold text-lg">
                        <option value="adult">Adult</option>
                        <option value="child">Child</option>
                    </select>
                    <div class="my-auto">
                        <img id="delete-${counter}" class="delete w-10 ml-2" src="/delete_red_icon.png" alt="delete_red_icon_image">
                    </div>
                </div>
            `;

            spawns?.insertAdjacentHTML("beforeend", profileDiv);
        };

        const handleDeleteClick = (e: MouseEvent) => {
            const target = e.target as HTMLElement;
            if (target && target.classList.contains("delete")) {
                const profileDiv = target.closest('div[id^="div-"]');
                if (profileDiv) {
                    profileDiv.remove();
                }
            }
        };

        add?.addEventListener("click", handleAddClick);
        spawns?.addEventListener("click", handleDeleteClick);

        // Cleanup event listeners on unmount
        return () => {
            add?.removeEventListener("click", handleAddClick);
            spawns?.removeEventListener("click", handleDeleteClick);
        };
    }, []);

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
                    <div className="flex">
                        <img
                            className="w-16"
                            src="/netflix_profile_image_yellow.jpg"
                            alt="netflix_profile_image_yellow_image"
                        />
                        <input
                            className="mx-4 border-2 border-gray-500 rounded px-3 w-full"
                            type="text"
                            placeholder="Profile name"
                        />
                        <select
                            name="ages"
                            id="ages"
                            className="bg-white font-bold text-lg"
                        >
                            <option value="adult">Adult</option>
                            <option value="child">Child</option>
                        </select>
                    </div>
                    <div id="spawns">
                        <div id="profileDiv" className="hidden flex my-2">
                            <img
                                className="w-16"
                                src="/netflix_profile_image_yellow.jpg"
                                alt="netflix_profile_image_yellow_image"
                            />
                            <input
                                className="mx-4 border-2 border-gray-500 rounded px-3 w-full"
                                type="text"
                                placeholder="Profile name"
                            />
                            <select
                                name="ages"
                                id="ages"
                                className="bg-white font-bold text-lg"
                            >
                                <option value="adult">Adult</option>
                                <option value="child">Child</option>
                            </select>
                            <div className="my-auto">
                                <img
                                    className="w-10 ml-2"
                                    src="/delete_red_icon.png"
                                    alt="delete_red_icon_image"
                                />
                            </div>
                        </div>
                    </div>
                    <button
                        className="bg-red-600 w-full my-4 text-white py-3 rounded hover:bg-red-700"
                    >SAVE
                    </button>
                </div>
            </section>

            <FooterSign />
        </main>
    );
};

export default CreateProfile;
