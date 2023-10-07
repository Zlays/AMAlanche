import React from 'react';
import './App.css';
import {Button, ChakraProvider, Stack} from '@chakra-ui/react'

function App() {
    return (
        <ChakraProvider>
            <div className="App">
                <h1>AMAlanche</h1>
                <Stack direction='row' spacing={4} align='center'>
                    <Button colorScheme='teal' variant='solid'>
                        Button
                    </Button>
                    <Button colorScheme='teal' variant='outline'>
                        Button
                    </Button>
                    <Button colorScheme='teal' variant='ghost'>
                        Button
                    </Button>
                    <Button colorScheme='teal' variant='link'>
                        Button
                    </Button>
                </Stack>
            </div>
        </ChakraProvider>
    );
}

export default App;
