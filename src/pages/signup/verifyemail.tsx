import { useEffect } from "react";
import FooterSign from "../../components/FooterSign";
import NavbarSign from "../../components/NavbarSign";

const VerifyEmail = () => {
    var email: any = new URLSearchParams(window.location.search).get("email");
    var password: any = new URLSearchParams(window.location.search).get("password");

    useEffect(() => {
        const emailSpan = document.getElementById("emailSpan")

        if (emailSpan) {
            emailSpan.innerHTML = email;
        }

        const hhref = document.getElementById("href")

        if (hhref) {
            hhref.addEventListener("click", () => {
                window.location.href = `/signup?email=${email}&password=${password}`;
            });
        }

        return () => {
            if (hhref) {
                hhref.removeEventListener("click", () => {
                    window.location.href = `/signup?email=${email}&password=${password}`;
                });
            }
            
            if (emailSpan) {
                emailSpan.innerHTML = email;
            }
        }
    }, []);

    return (
        <main>
            <NavbarSign />

            <section className="flex items-center justify-center text-center">
                <div className="py-32 w-[35%]">
                    <img
                        src="/icons8-aprobacion-50.png"
                        alt="aprobacion_image"
                        className="mx-auto mb-4"
                    />
                    <h6 className="text-sm">
                        STEP <span className="font-bold">2</span> OF <span className="font-bold"
                        >4</span>
                    </h6>
                    <h1 className="font-extrabold text-4xl mb-3">
                        Excelent! Now Let's verify your email
                    </h1>
                    <h4 className="text-gray-600 text-lg">
                        Click on the url we sent to <span
                            className="font-bold"
                            id="emailSpan"></span> to complete your verification.
                    </h4>
                    <h4 className="text-gray-600 text-lg">
                        When you verify your email, you'll can upgrade the security of
                        your account and get important notifications from Netflix.
                    </h4>
                    <div className="flex w-full mt-4">
                        <button
                            id="href"
                            className="bg-gray-300 hover:bg-gray-400 w-full rounded text-xl font-bold py-3"
                        >OMITIR</button>
                    </div>
                </div>
            </section>

            <FooterSign />
        </main>
    )
}

export default VerifyEmail;
