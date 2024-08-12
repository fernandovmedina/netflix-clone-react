import React, { useEffect } from "react";

const AddProfile = () => {
    useEffect(() => {
        const check = document.getElementById("check") as HTMLInputElement | null;
        const element = document.getElementById("checkH");

        if (element && check) {
            element.addEventListener("click", () => {
                check.checked = !check.checked;
            });
        }

        const continueBTN = document.getElementById("continue");
        const cancelBTN = document.getElementById("cancel");

        if (continueBTN) {
            continueBTN.addEventListener("click", () => {
                window.location.href = "/browse";
            });
        }

        if (cancelBTN) {
            cancelBTN.addEventListener("click", () => {
                window.location.href = "/browse";
            });
        }

        // Cleanup event listeners when component unmounts
        return () => {
            if (element && check) {
                element.removeEventListener("click", () => {
                    check.checked = !check.checked;
                });
            }

            if (continueBTN) {
                continueBTN.removeEventListener("click", () => {
                    window.location.href = "/browse";
                });
            }

            if (cancelBTN) {
                cancelBTN.removeEventListener("click", () => {
                    window.location.href = "/browse";
                });
            }
        };
    }, []);

    return (
        <main>
            <section
                className="bg-gradient-to-b from-black to-gray-950 text-white h-[100vh] flex flex-col items-center justify-center"
            >
                <div>
                    <h1 className="font-bold text-4xl">Add profile</h1>
                    <h4 className="text-gray-500 my-3">
                        Add a new profile for another person who watches Netflix.
                    </h4>
                    <div className="border-y-2 border-gray-500 flex items-center py-4">
                        <img
                            className="w-24"
                            src="https://occ-0-7553-114.1.nflxso.net/dnm/api/v6/vN7bi_My87NPKvsBoib006Llxzg/AAAABZfhNS88u5ao0M3F5X4HRBGCFsqdb2nncDt32YQHoM-1BeLJq93H30hWyleqclSwt1jNGm6l0tkeefKiiCOLLL5gNjpSjS_Xlaij.png?r=bd7"
                            alt="image_url"
                        />
                        <div>
                            <input
                                className="mx-2 bg-gray-600 text-white p-2"
                                type="text"
                                placeholder="Name"
                                required
                            />
                        </div>
                        <input type="checkbox" className="mx-2" id="check" />
                        <h3 className="text-lg" id="checkH">Children?</h3>
                    </div>
                    <div className="mt-4">
                        <button
                            id="continue"
                            className="mr-2 border-2 border-gray-600 text-gray-600 hover:bg-red-600 hover:border-transparent hover:text-white font-semibold px-4 py-2"
                        >
                            Continue
                        </button>
                        <button
                            id="cancel"
                            className="ml-2 border-2 border-gray-600 text-gray-600 hover:border-white hover:text-white px-4 py-2 font-semibold"
                        >
                            Cancel
                        </button>
                    </div>
                </div>
            </section>
        </main>
    );
};

export default AddProfile;
