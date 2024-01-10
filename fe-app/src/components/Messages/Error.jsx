import React, { useEffect } from 'react'

export default function Error({errorMessage, error, setError}) {

    useEffect(() => {
        if(error === true) {
            setTimeout(() => {
                setError(false);
            }, 3000)
        }
    }, [error])

    return (
        <div>
            <p className='error'>{errorMessage}</p>
        </div>
    )
}
