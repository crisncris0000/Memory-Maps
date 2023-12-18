import React, { useEffect, useState } from 'react';
import Modal from 'react-bootstrap/Modal';
import axios from 'axios';

export default function MarkerInfo({ show, setShow, markerPost}) {

  const handleClose = () => setShow(false);

  useEffect(() => {
    if(show) {
      axios.get(`http://localhost:8080/marker-post/images/${markerPost.id}`)
        .then((response) => {
          console.log(response.data);
        }).catch((error) => {
          console.log(error);
        })
    }
  })

  return (
    <>
      {show &&
        <Modal show={show} onHide={handleClose} animation={false} className="location-modal mx-auto" centered>
          <Modal.Header> 
            <Modal.Title>
              <div className="img-container">
              </div>
            </Modal.Title>
          </Modal.Header>
                  
          <Modal.Body>
          
          </Modal.Body>
        </Modal>
      }
    </>
  )
}
