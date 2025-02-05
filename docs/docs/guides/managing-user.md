import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';
import CodeBlock from '@theme/CodeBlock';

# Managing Users

A project in Shield looks like

```json
{
    "users": [
        {
            "id": "598688c6-8c6d-487f-b324-ef3f4af120bb",
            "name": "John Doe",
            "slug": "",
            "email": "john.doe@odpf.io",
            "metadata": {
                "role": "\"user-1\""
            },
            "createdAt": "2022-12-09T10:45:19.134019Z",
            "updatedAt": "2022-12-09T10:45:19.134019Z"
        }
    ]
}
```

One thing to note here is that Shield only allow to have metadata key from a specific set of keys. This constraint is only for users. We can add metadata key using this [metadata key API](./adding-metadata-key)

## API Interface

### Create users

<Tabs groupId="api">
  <TabItem value="HTTP" label="HTTP" default>
        <CodeBlock className="language-bash">
    {`$ curl --location --request POST 'http://localhost:8000/admin/v1beta1/users'
--header 'Content-Type: application/json'
--header 'Accept: application/json'
--header 'X-Shield-Email: admin@odpf.io'
--data-raw '{
  "name": "Jonny Doe",
  "email": "jonny.doe@odpf.io",
  "metadata": {
      "role": "user-3"
  }
}'`}
    </CodeBlock>
  </TabItem>
  <TabItem value="CLI" label="CLI" default>
<CodeBlock>

`$ shield user create --file=user.yaml`
</CodeBlock>

  </TabItem>
</Tabs>

### List users

<Tabs groupId="api">
  <TabItem value="HTTP" label="HTTP" default>
        <CodeBlock className="language-bash">
    {`curl --location --request GET 'http://localhost:8000/admin/v1beta1/users'
--header 'Accept: application/json'`}
    </CodeBlock>
  </TabItem>
  <TabItem value="CLI" label="CLI" default>
<CodeBlock>

`$ shield user list`
</CodeBlock>

  </TabItem>
</Tabs>

### Get Users

<Tabs groupId="api">
  <TabItem value="HTTP" label="HTTP" default>
        <CodeBlock className="language-bash">
    {`$ curl --location --request GET 'http://localhost:8000/admin/v1beta1/users/e9fba4af-ab23-4631-abba-597b1c8e6608'
--header 'Accept: application/json''`}
    </CodeBlock>
  </TabItem>
  <TabItem value="CLI" label="CLI" default>
<CodeBlock>

`$ shield user view e9fba4af-ab23-4631-abba-597b1c8e6608 --metadata`
</CodeBlock>

  </TabItem>
</Tabs>

### Update Projects

<Tabs groupId="api">
  <TabItem value="HTTP" label="HTTP" default>
        <CodeBlock className="language-bash">
    {`$ curl --location --request PUT 'http://localhost:8000/admin/v1beta1/users/e9fba4af-ab23-4631-abba-597b1c8e6608'
--header 'Content-Type: application/json'
--header 'Accept: application/json'
--data-raw '{
  "name": "Jonny Doe",
  "email": "john.doe001@odpf.io",
  "metadata": {
      "role" :   "user-3"
  }
}'`}
    </CodeBlock>
  </TabItem>
  <TabItem value="CLI" label="CLI" default>
<CodeBlock>

`$ shield user edit e9fba4af-ab23-4631-abba-597b1c8e6608 --file=user.yaml`
</CodeBlock>

  </TabItem>
</Tabs>