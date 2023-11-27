import React, { useState } from 'react';
import Button from 'react-bootstrap/Button';
import Modal from 'react-bootstrap/Modal';
import axios from 'axios';

export default function LocationInfo({ show, setShow, longitude, latitude, onHide }) {

    const [image, setImage] = useState(null)
    const [description, setDescription] = useState(null)

    const handleImageChange = (e) => {
        const file = e.target.files[0]
        setImage(file)
    }

    const handleClose = () => {
        onHide()
        setShow(false);
    }

    const handleOnSubmit = () => {
    
        const formData = new FormData();
        formData.append("latitude", latitude);
        formData.append("longitude", longitude);
        formData.append("image", image);
        formData.append("description", description);
        formData.append("visibilityID", 1);
        formData.append("userID", 1);
    
        axios.post("http://localhost:8080/marker-posts/new", formData)
            .then((response) => {
                console.log(response.data);
            })
            .catch((error) => {
                console.log(error);
            });
    };
    

    return (
        <>
            <Modal show={show} onHide={handleClose} animation={false} className="location-modal mx-auto" centered>
                <Modal.Header closeButton> 
                    <Modal.Title>Marker Form</Modal.Title>
                </Modal.Header>
                <form onSubmit={handleOnSubmit}>
                    <Modal.Body>
                        <div className="form-group">
                            <label className="upload-label">Image upload *</label>
                            <input type="file" required className="form-control upload-input" onChange={handleImageChange} />
                        </div>

                        <div className="form-group">
                            <label className="description-label">Description *</label>
                            <textarea type="text" placeholder="Please enter description" className="form-control description" 
                            onChange={e => setDescription(e.target.value)}/>
                        </div>
                    </Modal.Body>
                    <Modal.Footer>
                        <Button variant="secondary" onClick={handleClose}>
                            Close
                        </Button>
                        <Button variant="primary" type="submit">
                            Submit
                        </Button>
                    
                    </Modal.Footer>
                </form>
            </Modal>
        </>
        
    );
}
