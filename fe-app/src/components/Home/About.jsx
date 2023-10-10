import React from 'react'

export default function About() {
  return (
    <>
        <div className="about-me">
            <h1 className="header">What's this website about?</h1>
            <p>
                This website is a place where users are able to place pinpoints on a map and within that pinpoint
                the user can upload photos of their time at that specific part of the globe.
            </p>
        </div>

        <div className="info-container">
          <div className="target-audience">
            <h1 className="header">Who is this for</h1>

            <p>
              Primarily this applciation is for those who love to travel to different places, and would like to take pictures
              to create family memories, or memories for yourself. If you are not within that category that is Ok! It can
              be used for summer vacations, visiting family, and much more!
            </p>
          </div>

          <div className="getting-started">
            <h1 className="header">Getting Started</h1>
            <p>To be Determined</p>
          </div>
        </div>
    </>
  )
}
