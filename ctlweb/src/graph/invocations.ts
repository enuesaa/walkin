import { gql, useQuery } from 'urql'
import { Query } from './types'

const query = gql`
  query {
    invocations {
      id
      status
      method
      url
      requestBody
      responseBody
    }
  }
`

export const useListInvocations = () => useQuery<Query>({ query })[0]
