import { useEffect } from "react";
import FooterSign from "../components/FooterSign";
import NavbarSign from "../components/NavbarSign";

const Signup = () => {
    var email = new URLSearchParams(window.location.search).get("email");
    var password = new URLSearchParams(window.location.search).get("password");

    useEffect(() => {
        document.getElementById("btn")!.addEventListener("click", () => {
            window.location.href = `/signup/planform?email=${email}&password=${password}`;
        });
    }, []);

    return (
        <main>
            <NavbarSign />

            <section className="flex items-center justify-center">
                <div className="text-center py-36 w-[30%]">
                    <img
                        className="mx-auto"
                        src="/correcto_circulo.png"
                        alt="correcto_circulo_image"
                    />
                    <h5 className="text-sm mt-4">
                        STEP <span className="font-bold">2</span> OF <span
                            className="font-bold">4</span>
                    </h5>
                    <h1 className="font-bold text-4xl mb-4">Select your plan</h1>
                    <div className="flex items-center my-2">
                        <img src="/correcto.png" alt="correcto_image" />
                        <h4 className="text-xl font-medium ml-2">
                            Without compromises, cancel whenever you want
                        </h4>
                    </div>
                    <div className="flex items-center my-2">
                        <img src="/correcto.png" alt="correcto_image" />
                        <h4 className="text-xl font-medium ml-2">
                            Entertainment without an end and low cost
                        </h4>
                    </div>
                    <div className="flex items-center my-2">
                        <img src="/correcto.png" alt="correcto_image" />
                        <h4 className="text-xl font-medium ml-2">
                            Enjoy Netflix on all your devices
                        </h4>
                    </div>
                    <div className="w-full flex">
                        <button
                            id="btn"
                            className="bg-red-600 w-full text-white hover:bg-red-700 rounded py-3 mt-4 mx-auto"
                        >NEXT</button>
                    </div>
                </div>
            </section>

            <FooterSign />
        </main>
    )
}

export default Signup;
