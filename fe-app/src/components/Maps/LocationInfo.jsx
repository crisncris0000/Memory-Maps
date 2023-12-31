import React, { useState } from 'react';
import Button from 'react-bootstrap/Button';
import Modal from 'react-bootstrap/Modal';
import axios from 'axios';
import { useSelector } from 'react-redux';
import Compressor from 'compressorjs';


export default function LocationInfo({ show, setShow, longitude, latitude, onHide }) {

    const user = useSelector((state) => state.user.value);
    const [imageData, setImageData] = useState([]);
    const [description, setDescription] = useState("");
    const [likes, setLikes] = useState(0);
    const [visibilityID, setVisibilityID] = useState(0);

    const handleClose = () => {
        onHide();
        setImageData([]);
        setShow(false);
    }

    const handleOnSubmit = (e) => {
    
        axios.post("http://localhost:8080/marker-posts/new", {
            latitude,
            longitude,
            description,
            imageData,
            likes,
            visibilityID,
            userEmail: user.email,
        })
            .then((response) => {
                console.log(response.data);
            })
            .catch((error) => {
                console.log(error);
            });
    };

    const handleImageChange = (event) => {
        const files = event.target.files;
    
        if (!files || files.length === 0) {
            return;
        }
    
        const newImages = [];
    
        for (const file of files) {
            new Compressor(file, {
                quality: 0.7,
                success(result) {
                    const reader = new FileReader();
                    reader.onloadend = () => {
                        const base64String = reader.result.split(',')[1];
    
                        const newImageData = {
                            image: base64String,
                            mimeType: result.type
                        };
    
                        newImages.push(newImageData);

                        if (newImages.length === files.length) {
                            setImageData((prevImageData) => [...prevImageData, ...newImages]);
                        }
                    };
                    reader.readAsDataURL(result);
                },
                error(err) {
                    console.error('[Compressor.js] Error:', err.message);
                },
            });
        }
    }
    

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
                            <input type="file" required className="form-control upload-input" multiple onChange={handleImageChange} />
                        </div>

                        <div className="form-group">
                            <label className="description-label">Description *</label>
                            <textarea type="text" placeholder="Please enter description" className="form-control description" 
                            onChange={e => setDescription(e.target.value)}/>
                        </div>


                        <div className="visibility-container">
                            <div className="radio-container">
                                <input type="radio" className="radio" name="visibility" id="public" onClick={() => {setVisibilityID(1)}} required/>
                                <label className="radio-label" htmlFor="public">Public</label>
                            </div>
                            <div className="radio-container">
                                <input type="radio" className="radio" name="visibility" id="private" onClick={() => {setVisibilityID(2)}}required/>
                                <label className="radio-label" htmlFor="private">Private</label>
                            </div>
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
