import React from 'react'

export default function PreviewImages({ imageData, setImageData }) {
  return (
    <div className='preview-image-container'>
      {imageData.map((img) => (
        <img src={`data:${img.mimeType};base64,${img.image}`}/>
      ))}
    </div>
  )
}
