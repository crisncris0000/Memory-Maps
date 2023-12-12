import React, { useEffect } from 'react';
import Modal from 'react-bootstrap/Modal';
import Map from '../../images/pinpoint-map.jpg';

export default function MarkerInfo({ show, setShow, markerPost}) {

  const handleClose = () => setShow(false);

  useEffect(() => {
    
  })

  return (
    <>
      {show &&
        <Modal show={show} onHide={handleClose} animation={false} className="location-modal mx-auto" centered>
          <Modal.Header> 
            <Modal.Title>
              <div className="img-container">
                <img src={Map}/>
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
