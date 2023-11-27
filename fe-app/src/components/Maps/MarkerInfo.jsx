import React, { useEffect } from 'react';
import Modal from 'react-bootstrap/Modal';

export default function MarkerInfo({ show, setShow, markerPost}) {

  const handleClose = () => setShow(false);

  useEffect(() => {
    
  })

  return (
    <>
      {show &&
        <Modal show={show} onHide={handleClose} animation={false} className="location-modal mx-auto" centered>
          <Modal.Header closeButton> 
            <Modal.Title>
            </Modal.Title>
          </Modal.Header>
                  
          <Modal.Body>
          
          </Modal.Body>
        </Modal>
      }
    </>
  )
}
