import React from 'react';
import disponible from './disponible';

function Page(){

    const Carte = () => (
        <div className="inline-block max-w-sm mx-auto bg-blue rounded overflow-hidden shadow-lg">
            {disponible() ? (<p className="text-gren text-right"> Disponible </p>) : (<p className="text-red text-right"> Non disponible </p>)}
            <div className="px-6 py-7">
                <div className="font-bold text-xl mb-2">201</div>
            </div>
            <div className="card-image-container">
                <img src="/image/essaie.png" alt="voici notre salle" className="w-full h-auto rounded-t" />
            </div>
        </div>
    );
    
    return (
        <main className="bg-celeste min-h-screen">
            <h1 className="text-5xl font-bold text-grayblue">
                <div className="p-12 flex items-center justify-center w-screen ">
                    Vérifier la disponibilité des salles !!
                </div>
            </h1>
            <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-5 gap-4">
                {[...Array(32)].map((_, index) => (
                    <Carte key={index} />
                ))}
            </div>
        </main>
    )
}

export default Page;
