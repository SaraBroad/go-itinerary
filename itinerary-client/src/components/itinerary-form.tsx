import React, { useState } from 'react'

export const ItineraryForm = () => {
    const [name, setName] = useState('')
    const [startDate, setStartDate] = useState('')
    const [endDate, setEndDate] = useState('')
    const [locations, setLocations] = useState<string[]>([])

    const handleLocationChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        const value = event.target.value
        setLocations([...locations, value])
    }

    return (
        <>
        <label>
            Name{': '}
            <input value={name} onChange={e => setName(e.target.value)} />
        </label>
        <label>
            StartDate{': '}
            <input value={startDate} onChange={e => setStartDate(e.target.value)} />
        </label>
        <label>
            End Date{': '}
            <input value={endDate} onChange={e => setEndDate(e.target.value)} />
        </label>
        <label>
            Locations{': '}
            <input value={locations} onChange={handleLocationChange} />
        </label>
        </>
    )
}