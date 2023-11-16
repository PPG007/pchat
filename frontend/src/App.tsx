import { Fragment } from 'react'
import { ExampleEnum, ExampleMessage } from "./pb/example.ts";

function App() {
  return (
    <Fragment>
      <button onClick={() => {
        const e = ExampleMessage.create();
        e.id = 1;
        e.status = ExampleEnum.FINISHED;
        e.name = 'PPG007';
        e.scores = {
          x: 1,
        }
        console.log(e)
      }}>click
      </button>
    </Fragment>
  )
}

export default App
