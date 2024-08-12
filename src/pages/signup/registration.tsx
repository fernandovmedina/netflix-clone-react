import NavbarSign from "../../components/NavbarSign";
import FooterSign from "../../components/FooterSign";

const Registration = () => {
    var btn: any = document.getElementById("btn");
    var email = new URLSearchParams(window.location.search).get("email");

    btn.addEventListener("click", () => {
        window.location.href = `/signup/regform?email=${email}`;
    });

    return (
        <main>
            <NavbarSign />

            <section className="flex items-center justify-center">
                <div className="py-40 w-[40%]">
                    <img src="/Devices.png" alt="devices_image" className="w-60 mx-auto" />
                    <p className="text-center mt-10 mb-2 text-sm">
                        STEP <span className="font-bold">1</span> OF <span className="font-bold"
                        >4</span>
                    </p>
                    <h2 className="text-center font-bold text-4xl">
                        Complete the account configuration
                    </h2>
                    <h4 className="text-center mt-2 text-xl">
                        Netflix is designed four you.
                    </h4>
                    <h4 className="text-center text-xl">
                        Create a password to start watching Netflix.
                    </h4>
                    <div className="w-full flex justify-center mt-4">
                        <button
                            id="btn"
                            className="bg-red-600 py-3 hover:bg-red-700 text-white w-full text-center rounded"
                        >NEXT</button
                        >
                    </div>
                </div>
            </section>

            <FooterSign />
        </main>
    )
}

export default Registration;
