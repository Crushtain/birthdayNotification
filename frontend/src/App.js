import { Container, Row, Col  } from "react-bootstrap"
import React, {useEffect, useState} from "react";
import axios from "axios";

function App() {
  const [apiData, setApiData] = useState(false);

  useEffect(() => {
      const fetchData = async () => {
          try {
              const apiUrl = "http://localhost:8080/";
              const response = await axios.get(apiUrl);

              if (response.status === 201) {
                  if (response?.data.status === "Ok")
                  setApiData(response?.data?.users);
              }
          } catch (error){
            console.log(error.response);
          }



      };
      fetchData();


      return () => {};
  }, []);

    console.log(apiData)


    return (
    <Container>
      <Row>
         <Col xs="12" className="py-2">
           <h1 className="text-center">
               React application with Go fiber backend
           </h1>
         </Col>
          {apiData &&
              apiData.map((users, index) => (
                  <Col xs="4" className="py-5">
                      <div>{users.username}</div>
                      <div>{users.birthday}</div>
                  </Col>
              ))}
      </Row>
    </Container>
  );
}

export default App;
