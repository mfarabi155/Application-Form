import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useNavigate, useParams } from 'react-router-dom';

function PersonForm() {
    const [name, setName] = useState("");
    const [identityNumberState, setIdentityNumber] = useState("");
    const [email, setEmail] = useState("");
    const [dateOfBirth, setDateOfBirth] = useState("");
    const navigate = useNavigate();
    const { identityNumber: paramIdentityNumber } = useParams();
    const isEditing = Boolean(paramIdentityNumber);

    useEffect(() => {
        if (isEditing) {
            axios.get(`http://localhost:8080/persons/${paramIdentityNumber}`)
                .then(response => {
                    const person = response.data;
                    setName(person.name);
                    setIdentityNumber(person.identity_number);
                    setEmail(person.email);
                    setDateOfBirth(person.date_of_birth);
                })
                .catch(error => {
                    console.error('Error fetching person data:', error);
                });
        }
    }, [paramIdentityNumber, isEditing]);

    const handleIdentityNumberChange = (e) => {
        const value = e.target.value;
        if (/^\d*$/.test(value)) {
            setIdentityNumber(value);
        }
    };

    const handleSubmit = (event) => {
        event.preventDefault();

        const person = {
            name,
            identity_number: identityNumberState,
            email,
            date_of_birth: dateOfBirth,
        };

        if (isEditing) {
            axios.put(`http://localhost:8080/persons/${paramIdentityNumber}`, person)
                .then(() => {
                    navigate("/");
                })
                .catch(error => {
                    console.error('Error updating person:', error);
                });
        } else {
            axios.post('http://localhost:8080/persons', person)
                .then(() => {
                    navigate("/");
                })
                .catch(error => {
                    console.error('Error adding person:', error);
                });
        }
    };

    return (
        <div className="container py-4">
            <h1>{isEditing ? 'Edit Person' : 'Add Person'}</h1>
            <form onSubmit={handleSubmit}>
                <div className="mb-3">
                    <label className="form-label">Name:</label>
                    <input type="text" className="form-control" value={name} onChange={(e) => setName(e.target.value)} required />
                </div>
                <div className="mb-3">
                    <label className="form-label">Identity Number:</label>
                    <input type="text" className="form-control" value={identityNumberState} onChange={handleIdentityNumberChange} required />
                </div>
                <div className="mb-3">
                    <label className="form-label">Email:</label>
                    <input type="email" className="form-control" value={email} onChange={(e) => setEmail(e.target.value)} required />
                </div>
                <div className="mb-3">
                    <label className="form-label">Date of Birth:</label>
                    <input type="date" className="form-control" value={dateOfBirth} onChange={(e) => setDateOfBirth(e.target.value)} required />
                </div>
                <button type="submit" className="btn btn-primary">{isEditing ? 'Update' : 'Submit'}</button>
            </form>
        </div>
    );
}

export default PersonForm;
