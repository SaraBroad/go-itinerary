import React, { useState } from 'react'

export const ItineraryForm = () => {
    const [name, setName] = useState('')
    const [startDate, setStartDate] = useState('')
    const [endDate, setEndDate] = useState('')
    const [locations, setLocations] = useState([])
    return (
        <>
        <label>
            Name{': '}
            <input value={name} onChange={e => setName(e.target.value)} />
        </label>
        </>
    )
}