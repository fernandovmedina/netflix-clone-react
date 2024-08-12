import FooterSign from "../../components/FooterSign";
import NavbarSign from "../../components/NavbarSign";

const Planform = () => {
    var nextBTN: any = document.getElementById("nextBTN");
    var email = new URLSearchParams(window.location.search).get("email");
    var password = new URLSearchParams(window.location.search).get("password");

    nextBTN.addEventListener("click", () => {
        window.location.href = `/signup/paymentPicker?email=${email}&password=${password}&plan=${GRGH()}`;
    });

    var div_1: any = document.getElementById("div_1");
    var div_2: any = document.getElementById("div_2");
    var div_3: any = document.getElementById("div_3");

    var rgh_1: any = document.getElementById("rgh_1");
    var rgh_2: any = document.getElementById("rgh_2");
    var rgh_3: any = document.getElementById("rgh_3");

    var bool_1: boolean = false;
    var bool_2: boolean = false;
    var bool_3: boolean = false;

    div_1.addEventListener("click", () => {
        bool_1 = !bool_1;
        if (bool_2 || bool_3) {
            bool_2 = false;
            bool_3 = false;
            rgh_2.style.display = "none";
            rgh_3.style.display = "none";
        }
        rgh_1.style.display = "block";
    });

    div_2.addEventListener("click", () => {
        bool_2 = !bool_2;
        if (bool_1 || bool_3) {
            bool_1 = false;
            bool_3 = false;
            rgh_1.style.display = "none";
            rgh_3.style.display = "none";
        }
        rgh_2.style.display = "block";
    });

    div_3.addEventListener("click", () => {
        bool_3 = !bool_3;
        if (bool_1 || bool_2) {
            bool_1 = false;
            bool_2 = false;
            rgh_1.style.display = "none";
            rgh_2.style.display = "none";
        }
        rgh_3.style.display = "block";
    });

    function GRGH() {
        if (bool_1) {
            return "premium";
        } else if (bool_2) {
            return "standard";
        } else {
            return "ads";
        }
    }

    return (
        <main>
            <NavbarSign />

            <section>
                <div className="py-12 px-40">
                    <h5 className="text-sm">
                        STEP <span className="font-bold">3</span> OF <span className="font-bold"
                        >4</span>
                    </h5>
                    <h1 className="font-bold text-3xl">Choose the ideal plan for you</h1>
                    <div className="w-full flex mt-4">
                        <div
                            className="w-1/3 border-2 border-gray-300 rounded py-2 px-2 mx-1"
                            id="div_1"
                        >
                            <div
                                className="bg-gradient-to-br from-indigo-800 via-violet-900 to-pink-600 text-white rounded p-3"
                            >
                                <h2 className="font-bold text-xl">Premium</h2>
                                <h2 className="font-bold text-sm">4K + HDR</h2>
                                <div className="hidden pt-2" id="rgh_1">
                                    <img
                                        src="/correcto_blanco.png"
                                        alt="correcto_blanco_image"
                                        className="w-5"
                                    />
                                </div>
                            </div>
                            <div className="px-3 py-4">
                                <div className="border-b-2 border-gray-300 py-3">
                                    <h2 className="text-gray-600 text-sm">Monthly price</h2>
                                    <h2 className="font-bold">$299</h2>
                                </div>
                                <div className="border-b-2 border-gray-300 py-3">
                                    <h2 className="text-gray-600 text-sm">
                                        Audio and video quality
                                    </h2>
                                    <h2 className="font-bold">Optimal</h2>
                                </div>
                                <div className="border-b-2 border-gray-300 py-3">
                                    <h2 className="text-gray-600 text-sm">Resolution</h2>
                                    <h2 className="font-bold">4K (Ultra HD) + HDR</h2>
                                </div>
                                <div className="border-b-2 border-gray-300 py-3">
                                    <h2 className="text-gray-600 text-sm">
                                        Spatial audio (immersive sound)
                                    </h2>
                                    <h2 className="font-bold">Included</h2>
                                </div>
                                <div className="border-b-2 border-gray-300 py-3">
                                    <h2 className="text-gray-600 text-sm">
                                        Compatible devices
                                    </h2>
                                    <h2 className="font-bold">TV, pc, telephone, tablet</h2>
                                </div>
                                <div className="border-b-2 border-gray-300 py-3">
                                    <h2 className="text-gray-600 text-sm">
                                        Home devices which you can see at the same time
                                    </h2>
                                    <h2 className="font-bold">4</h2>
                                </div>
                                <div className="border-b-2 border-gray-300 py-3">
                                    <h2 className="text-gray-600 text-sm">
                                        Download devices
                                    </h2>
                                    <h2 className="font-bold">6</h2>
                                </div>
                                <div className="py-3">
                                    <h2 className="text-gray-600 text-sm">Adds</h2>
                                    <h2 className="font-bold">No Adds</h2>
                                </div>
                            </div>
                        </div>
                        <div
                            className="w-1/3 border-2 border-gray-300 rounded py-2 px-2 mx-1"
                            id="div_2"
                        >
                            <div
                                className="bg-gradient-to-br from-blue-800 via-indigo-700 to-purple-800 text-white rounded p-3"
                            >
                                <h2 className="font-bold text-xl">Standard</h2>
                                <h2 className="font-bold text-sm">1080p</h2>
                                <div className="hidden pt-2" id="rgh_2">
                                    <img
                                        src="/correcto_blanco.png"
                                        alt="correcto_blanco_image"
                                        className="w-5"
                                    />
                                </div>
                            </div>
                            <div className="px-3 py-4">
                                <div className="border-b-2 border-gray-300 py-3">
                                    <h2 className="text-gray-600 text-sm">Monthly price</h2>
                                    <h2 className="font-bold">$219</h2>
                                </div>
                                <div className="border-b-2 border-gray-300 py-3">
                                    <h2 className="text-gray-600 text-sm">
                                        Audio and video quality
                                    </h2>
                                    <h2 className="font-bold">Good</h2>
                                </div>
                                <div className="border-b-2 border-gray-300 py-3">
                                    <h2 className="text-gray-600 text-sm">Resolution</h2>
                                    <h2 className="font-bold">1080p (Full HD)</h2>
                                </div>
                                <div className="border-b-2 border-gray-300 py-3">
                                    <h2 className="text-gray-600 text-sm">
                                        Compatible devices
                                    </h2>
                                    <h2 className="font-bold">TV, pc, telephone, tablet</h2>
                                </div>
                                <div className="border-b-2 border-gray-300 py-3">
                                    <h2 className="text-gray-600 text-sm">
                                        Home devices which you can see at the same time
                                    </h2>
                                    <h2 className="font-bold">2</h2>
                                </div>
                                <div className="border-b-2 border-gray-300 py-3">
                                    <h2 className="text-gray-600 text-sm">
                                        Download devices
                                    </h2>
                                    <h2 className="font-bold">2</h2>
                                </div>
                                <div className="py-3">
                                    <h2 className="text-gray-600 text-sm">Adds</h2>
                                    <h2 className="font-bold">No Adds</h2>
                                </div>
                            </div>
                        </div>
                        <div
                            className="w-1/3 border-2 border-gray-300 rounded py-2 px-2 mx-1"
                            id="div_3"
                        >
                            <div
                                className="bg-gradient-to-br from-blue-950 via-indigo-700 to-blue-700 text-white rounded p-3"
                            >
                                <h2 className="font-bold text-xl">Standard with adds</h2>
                                <h2 className="font-bold text-sm">1080p</h2>
                                <div className="hidden pt-2" id="rgh_3">
                                    <img
                                        src="/correcto_blanco.png"
                                        alt="correcto_blanco_image"
                                        className="w-5"
                                    />
                                </div>
                            </div>
                            <div className="px-3 py-4">
                                <div className="border-b-2 border-gray-300 py-3">
                                    <h2 className="text-gray-600 text-sm">Monthly price</h2>
                                    <h2 className="font-bold">$99</h2>
                                </div>
                                <div className="border-b-2 border-gray-300 py-3">
                                    <h2 className="text-gray-600 text-sm">
                                        Audio and video quality
                                    </h2>
                                    <h2 className="font-bold">Good</h2>
                                </div>
                                <div className="border-b-2 border-gray-300 py-3">
                                    <h2 className="text-gray-600 text-sm">Resolution</h2>
                                    <h2 className="font-bold">1080p (Full HD)</h2>
                                </div>
                                <div className="border-b-2 border-gray-300 py-3">
                                    <h2 className="text-gray-600 text-sm">
                                        Compatible devices
                                    </h2>
                                    <h2 className="font-bold">TV, pc, telephone, tablet</h2>
                                </div>
                                <div className="border-b-2 border-gray-300 py-3">
                                    <h2 className="text-gray-600 text-sm">
                                        Home devices which you can see at the same time
                                    </h2>
                                    <h2 className="font-bold">2</h2>
                                </div>
                                <div className="border-b-2 border-gray-300 py-3">
                                    <h2 className="text-gray-600 text-sm">
                                        Download devices
                                    </h2>
                                    <h2 className="font-bold">2</h2>
                                </div>
                                <div className="py-3">
                                    <h2 className="text-gray-600 text-sm">Adds</h2>
                                    <h2 className="font-bold">Less than you can count</h2>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </section>

            <section className="px-40 mb-10 text-sm text-gray-500">
                <p className="mb-2">
                    <a
                        className="text-blue-600 hover:underline hover:underline-offset-1"
                        href="https://help.netflix.com/node/126831"
                    >More info about the plan with ads.</a
                    > If you select a plan with ads, you will be asked to provide your date
                    of birth for ad personalization and for other purposes consistent with
                    the Netflix <a
                        href="https://help.netflix.com/legal/privacy"
                        className="text-blue-600 hover:underline hover:underline-offset-1"
                    >Privacy Statement.</a>
                </p>
                <p className="mb-2">
                    The availability of content in Full HD (1080p), Ultra HD (4K) and
                    HDR depends on your internet service and the device in use. Not all
                    content is available in all resolutions. See <a
                        href="https://help.netflix.com/legal/termsofuse"
                        className="text-blue-600 hover:underline hover:underline-offset-1"
                    >Terms of Use</a
                    > for more information.
                </p>
                <p>
                    Only people who live with you can use your account. Add 1 extra
                    member with the Standard plan or up to 2 with the Premium plan. <a
                        className="text-blue-600 hover:underline hover:underline-offset-1"
                        href="https://help.netflix.com/pricing">More info.</a
                    > You can watch Netflix on 4 devices at the same time with the Premium
                    plan and on 2 with the Standard or Standard plan with ads.
                </p>
            </section>

            <div className="text-center mb-10">
                <button
                    id="nextBTN"
                    className="bg-red-600 text-white py-3 px-40 rounded hover:bg-red-700"
                >NEXT</button
                >
            </div>

            <FooterSign />
        </main>
    )
}

export default Planform;
