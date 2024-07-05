# What is Lunar?

Lunar is an API service for storing and retriving data or "Entities". An Entity is a collection of arbitrary key-value
pairs, that conform to a schema defined by the user.

Entities are organized into "Collections", which are essentially like a table in a database. Each collection has a unique
name, and defines "EntityTypes" which are the schemas that entities in that collection must conform to.

EntitySchemas are defined with a standard JSON schema format that is used to validate the data that is stored in the
collection. Each Collection defines a set of EntityTypes, and these EntityTypes can have different schemas and be related
to other EntityTypes in the collection in arbitrary ways.

EntityTypes are themselves defined with a JSON schema and provide a way to extend the built-in JSON schema format to allow
for more complex data structures and relationships between entities. For example, an EntityType can define a field that
extends the Object type to allow for nested entities, or a field that extends the Array type to allow for a list of
entities.

Lunar also has a concept of "EntityStates" which can be used to define versions of an Entity that have different values
for the same fields. This allows for an Entity to be modified in one state without affecting the values of the fields in
other states. For example, an Entity could have a "draft" state that is used for editing, and a "published" state that is
used for displaying the data to users. APIs are provided to manage the state of an Entity or a set of Entities, and to
query for Entities based on their state.

Collections of Entities are indexed in OpenSearch, which provides a fast and efficient way to search for Entities based on
their field values, as well as searching based on their State or EntityType. The search engine is distributed and can be
sharded across multiple nodes to provide fast and efficient search.

## What is Lunar used for?

Lunar is intended to be used as a backend service for applications that need to store and retrieve data in a flexible and
performant way. Often when building applications, the data model is not known ahead of time, and the ability to store
arbitrary data is required. Lunar provides a way to store this data in a structured way, while still allowing for
flexibility in the schema.

Example use cases for Lunar include:

- product data for an e-commerce website
- content for a content management system
- user data for a social media platform
- sensor data for an IoT application

A single Lunar instance can be used to store data for multiple applications, each with its own set of collections and
entity types. This allows for a high degree of flexibility in how data is stored and organized, while still providing a
consistent API for interacting with the data.

## Why would I not use a traditional database directly?

Traditional databases like MySQL, PostgreSQL, or MongoDB are great for storing structured data, but they can be limiting
when it comes to storing arbitrary data. For example, if you have an e-commerce website and you want to store product data,
adding a new field to the product schema can be difficult and require a migration of the database schema.

Lunar is designed around the concept of storing data with dynamically defined schemas, along with automatic validation and
migration of the data. This makes it easy to add new fields to the schema, or change the schema entirely, without having to
perform a database schema migration to the underlying database.

## Permissions Model

Lunar has a flexible permissions model that allows for fine-grained control over who can access and modify data in a
collection. Permissions are defined at both the collection level and the entity level, and can be set to allow or deny
access to specific operations. These permissions are dynamically defined and can be changed at any time, allowing for
flexible access control based on the needs of the application.

# Example Application

To demonstrate how Lunar can be used in a real-world application, we will build a simple content management system (CMS)
that allows users to create, edit, and publish articles. The CMS will have the following features:

- Users can create articles with a title, body, and author
- Users can edit articles and save them as drafts
- Users can publish articles to make them visible to other users
- Users can search for articles based on their title, author, or content

The CMS will use Lunar to store the articles and provide an API for interacting with the data. The Lunar instance will
have a single collection called "articles" that defines an EntityType for articles. The EntityType will have fields for the
title, body, author, and state of the article, as well as a field for the creation date.

The CMS will use the Lunar API to create, edit, publish, and search for articles, as well as manage the state of the
articles. The Lunar API will provide endpoints for these operations, as well as endpoints for managing the permissions of
the collection.
