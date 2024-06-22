import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Link, useNavigate } from 'react-router-dom';

function PersonTable() {
    const [persons, setPersons] = useState([]);
    const navigate = useNavigate();

    useEffect(() => {
        fetchData();
    }, []);

    const fetchData = () => {
        axios.get('http://localhost:8080/persons')
            .then(response => {
                setPersons(response.data);
            })
            .catch(error => {
                console.error('Error fetching data:', error);
            });
    };

    const handleDeleteClick = (identityNumber) => {
        axios.delete(`http://localhost:8080/persons/${identityNumber}`)
            .then(response => {
                fetchData();
            })
            .catch(error => {
                console.error('Error deleting data:', error);
            });
    };

    const handleEditClick = (identityNumber) => {
        navigate(`/edit-person/${identityNumber}`);
    };

    return (
        <div className="container py-4">
            <h1 className="display-4">Data Persons</h1>
            <table className="table table-striped">
                <thead>
                    <tr>
                        <th scope="col">No</th>
                        <th scope="col">Name</th>
                        <th scope="col">Identity Number</th>
                        <th scope="col">Email</th>
                        <th scope="col">Date of Birth</th>
                        <th scope="col">Actions</th>
                    </tr>
                </thead>
                <tbody>
                    {persons.map((person, index) => (
                        <tr key={person.ID}>
                            <td>{index + 1}</td>
                            <td>{person.name}</td>
                            <td>{person.identity_number}</td>
                            <td>{person.email}</td>
                            <td>{person.date_of_birth}</td>
                            <td>
                                <button
                                    className="btn btn-primary me-2"
                                    onClick={() => handleEditClick(person.identity_number)}>
                                    Edit
                                </button>
                                <button
                                    className="btn btn-danger"
                                    onClick={() => handleDeleteClick(person.identity_number)}>
                                    Delete
                                </button>
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>
            <Link to="/add-person">
                <button className="btn btn-primary">Add Person</button>
            </Link>
        </div>
    );
}

export default PersonTable;
