import FooterSign from "../../components/FooterSign";
import NavbarSign from "../../components/NavbarSign";

const Regform = () => {
    var email: any = new URLSearchParams(window.location.search).get("email");
    var emailInput: any = document.getElementById("email");
    var password: any = document.getElementById("password");

    emailInput.value = email;

    var btn: any = document.getElementById("btn");

    btn.addEventListener("click", () => {
        window.location.href = `/signup/verifyemail?email=${email}&password=${password.value}`;
    });

    return (
        <main>
            <NavbarSign />

            <section className="flex items-center justify-center">
                <div className="py-40 w-[35%]">
                    <h6 className="text-sm">
                        STEP <span className="font-bold">1</span> OF <span className="font-bold"
                        >4</span
                        >
                    </h6>
                    <h1 className="font-extrabold text-4xl pb-3">
                        Create a password to start your membership
                    </h1>
                    <h5 className="text-lg">Some steps more and that's it!</h5>
                    <h5 className="text-lg">We don't like procedures either.</h5>
                    <div className="flex flex-col mt-4">
                        <input
                            type="email"
                            name="email"
                            id="email"
                            className="px-3 py-2 border-2 border-gray-500 rounded"
                        />
                        <input
                            type="password"
                            name="password"
                            id="password"
                            placeholder="Password"
                            className="px-3 py-2 border-2 border-gray-500 rounded mt-2"
                        />
                        <button
                            id="btn"
                            className="text-center bg-red-600 mt-3 py-2 text-white rounded hover:bg-red-700"
                        >NEXT</button
                        >
                    </div>
                </div>
            </section>

            <FooterSign />
        </main>
    )
}

export default Regform;
