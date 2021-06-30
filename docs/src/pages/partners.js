
import React from 'react'
import Layout from '@theme/Layout'
import useBaseUrl from '@docusaurus/useBaseUrl'
import useDocusaurusContext from '@docusaurus/useDocusaurusContext'

const Avatar = ({ name, tags, image, description, link }) => {
  const context = useDocusaurusContext()
  const { siteConfig = {} } = context

  return (
    <div className="col col--4 margin-bottom--lg">
      <div
        className={'card'}
        style={{
          textAlign: 'center',
          backgroundColor: '#303846',
          color: '#ffffff',
        }}
      >
        <div
          className="card__image"
          style={{
            marginTop: 10,
            minHeight: 100,
            display: 'flex',
            alignSelf: 'center',
          }}
        >
          <img
            style={{ maxHeight: 100, objectFit: 'contain' }}
            src={image}
            alt={name}
          />
        </div>
        <div className="card__body">
          <div className="avatar">
            <div className="avatar__intro margin-left--none">
              <h4 className="avatar__name">{name}</h4>
              <ul style={{ padding: 0, marginTop: 10 }}>
                {tags.map((tag, index) => (
                  <li
                    className="badge badge--secondary"
                    key={index}
                    style={{ margin: 2 }}
                  >
                    {tag}
                  </li>
                ))}
              </ul>
              <small className="avatar__subtitle">{description}</small>
            </div>
          </div>
        </div>
        <div className="card__footer">
          <a href={link} target="_blank">
            <button className="button button--secondary">
              {siteConfig.themeConfig.t.global.discover}
            </button>
          </a>
        </div>
      </div>
    </div>
  )
}

function Partners() {
  const imgPath = useBaseUrl('/img/partners/')

  const context = useDocusaurusContext()
  const { siteConfig = {} } = context

  const founders = [
    {
      name: <>Cyberport Hong Kong</>,
      link: 'https://www.cyberport.hk/en',
      tags: ['Incubator'],
      image: imgPath + 'cyberport.jpg',
      description:""
    },
    {
      name: <>ETH Library Lab</>,
      link: 'https://www.librarylab.ethz.ch/',
      tags: ['Higher Education Institute'],
      image: imgPath + 'eth-library-lab.svg',
      description:""
    },
  ]

  const collaborators = [
  ]

  const patrons = [
    {
      name: 'CConstance',
      link: 'https://github.com/CConstance',
    },
    {
      name: 'yemingfeng',
      link: 'https://github.com/yemingfeng',
    },
  ]
  return (
    <Layout title={siteConfig.themeConfig.t.pages.partners.partnersLabel}>
      <div className="container margin-vert--lg">
        <h1 className="text--center margin-bottom--xl">
          {siteConfig.themeConfig.t.pages.partners.partnersLabel}
        </h1>
        <section>
          <h2>Founding Partners</h2>
          <div className="row">
            {founders.map((founder, index) => {
              const item = {
                ...founder,
                description:
                  siteConfig.themeConfig.t.pages.partners.founders[index]
                    .description,
              }
              return <Avatar {...item} key={index} />
            })}
          </div>
        </section>
        <section>
          <h2>{siteConfig.themeConfig.t.pages.partners.collaboratorsLabel}</h2>
          <div className="row">
            {collaborators.map((collaborator, index) => {
              const item = {
                ...collaborator,
                description:
                  siteConfig.themeConfig.t.pages.partners.collaborators[index]
                    .description,
              }
              return <Avatar {...item} key={index} />
            })}
          </div>
        </section>
        <section>
          <h2>{siteConfig.themeConfig.t.pages.partners.patronsLabel}</h2>
          <ul>
            {patrons.map((patron) => (
              <li>{patron.name}</li>
            ))}
          </ul>
        </section>
      </div>
    </Layout>
  )
}

export default Partners