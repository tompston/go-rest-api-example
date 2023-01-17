import * as F from './backend/public/gomarvin.gen.ts'
import { placeholderUser, placeholderUserSecond, PlaceholderUser } from './backend/public/placeholder.ts';
import { faker } from "https://cdn.skypack.dev/@faker-js/faker";

/**
 * Use the default API client to seed the db
 */
const client = F.defaultClient

async function seedFakeUsers(AMOUNT: number) {
    for (let i = 0; i < AMOUNT; i++) {

        const username = faker.internet.userName()

        const res = await F.UserEndpoints.RegisterUser(client, {
            password: faker.internet.password(),
            username: faker.internet.userName(username),
            email: faker.internet.email(username),
        });

        const data = await res.json();
        console.log(data.status);
    }
}


/**
 * Create a predefine user so that you can log in with credentials
 */
async function CreatePlaceholderUser(user: PlaceholderUser) {
    const res = await F.UserEndpoints.RegisterUser(client, {
        password: user.password,
        username: user.username,
        email: user.email,
    });
    const data = await res.json();
    console.log(data);
}

/**
 * Log in the predefined user after registration
 */
async function LoginPlaceholderUser(user: PlaceholderUser) {
    const res = await F.UserEndpoints.LoginUser(client, {
        password: user.password,
        username: user.username,
    });
    const data = await res.json();
    console.log(data);
    // console.log(data.data.token.access_token);
}

async function createTransactionBetweenAccounts(body: F.CreateTransactionBody) {
    const res = await F.TransactionEndpoints.CreateTransaction(client, body)
    const data = await res.json();
    console.log(data);
}


/** Create 100 fake users */
// seedFakeUsers(100)

/** Create 2 predefined placeholder users */
// CreatePlaceholderUser(placeholderUser)
// CreatePlaceholderUser(placeholderUserSecond)

/** Test if Login Works */
LoginPlaceholderUser(placeholderUser)
// LoginPlaceholderUser(placeholderUserSecond)

const user_1 = "2d74b17c-2041-49fd-861b-72cd1bc7a903"
const user_2 = "cc720122-04df-47c3-98ab-854bdedb9f8c"

async function seedTransactionsBetweenUsers(sender: string, reciever: string, AMOUNT: number) {
    for (let i = 0; i < AMOUNT; i++) {
        const amount = Math.random() * (100 - 1) + 1;
        await createTransactionBetweenAccounts(
            {
                sender_id: sender,
                reciever_id: reciever,
                amount: Math.round(amount) // TODO : Fix this, only ints are correct currently
            }
        )
    }
}

/** Seed fake transactions between both predefined users */
// seedTransactionsBetweenUsers(user_1, user_2, 50)
// seedTransactionsBetweenUsers(user_2, user_1, 50)