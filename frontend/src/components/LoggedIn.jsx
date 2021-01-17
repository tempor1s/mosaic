import React from 'react';
import { useAuth0 } from '../auth';
import { Button, Spinner, Text } from '@chakra-ui/react';

// TODO: the user is logged in in this example,
// so we should be able to fetch the users screenshots from the API
const LoggedIn = () => {
  const { loading, user, logout, isAuthenticated } = useAuth0();

  if (loading || !user) {
    return <Spinner />;
  }

  return (
    <>
      <Text fontWeight="bold" fontSize="3xl">
        Your screenshots
      </Text>
      <Text fontSize="xl">Hello, {user.email}</Text>
      {isAuthenticated && (
        // TODO: logout
        <Button onClick={() => logout()}>Logout</Button>
      )}
    </>
  );
};

export default LoggedIn;