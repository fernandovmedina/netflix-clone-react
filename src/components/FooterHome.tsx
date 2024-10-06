import React, { useEffect, useState } from 'react';

const FooterHome = () => {
    const [currentYear, setCurrentYear] = useState<number>(new Date().getFullYear());

    useEffect(() => {
        setCurrentYear(new Date().getFullYear());
    }, []);

    return (
        <footer className="bg-black w-full text-gray-400 text-sm px-44">
            <div className="flex flex-row">
                <img className='w-6 mx-3' src="/facebook_white.png" alt="facebook logo" />
                <img className='w-6 mx-3' src="/instagram_white.png" alt="instagram logo" />
                <img className='w-6 mx-3' src="/x_white.png" alt="x logo" />
                <img className='w-6 mx-3' src="/youtube_white.png" alt="youtube logo" />
            </div>
            
            <div className="w-full flex my-4">
                <div className='w-1/4 flex flex-col'>
                    <a href="#" className='my-1 hover:underline hover:underline-offset-1'>Descriptive audio</a>
                    <a href="#" className='my-1 hover:underline hover:underline-offset-1'>Investor's relationship</a>
                    <a href="#" className='my-1 hover:underline hover:underline-offset-1'>Privacy</a>
                    <a href="#" className='my-1 hover:underline hover:underline-offset-1'>Contact us</a>
                </div>
                <div className='w-1/4 flex flex-col'>
                    <a href="#" className='py-1 hover:underline hover:underline-offset-1'>Help center</a>
                    <a href="#" className='py-1 hover:underline hover:underline-offset-1'>Jobs</a>
                    <a href="#" className='py-1 hover:underline hover:underline-offset-1'>Legal advices</a>
                </div>
                <div className='w-1/4 flex flex-col'>
                    <a href="#" className='py-1 hover:underline hover:underline-offset-1'>Gift cards</a>
                    <a href="#" className='py-1 hover:underline hover:underline-offset-1'>Netflix store</a>
                    <a href="#" className='py-1 hover:underline hover:underline-offset-1'>Cookies preferences</a>
                </div>
                <div className='w-1/4 flex flex-col'>
                    <a href="#" className='py-1 hover:underline hover:underline-offset-1'>Press</a>
                    <a href="#" className='py-1 hover:underline hover:underline-offset-1'>Termins of use</a>
                    <a href="#" className='py-1 hover:underline hover:underline-offset-1'>Corporative information</a>
                </div>
            </div>

            <div>
                <button className='border border-gray-400 text-sm px-3 py-1 text-gray-400 hover:text-white'>Service code</button>
            </div>

            <div className='py-3'>
                <h5 className='text-xs text-gray-400'>1997 - {currentYear} Netflix, Inc</h5>
            </div>
        </footer>
    );
}

export default FooterHome;
