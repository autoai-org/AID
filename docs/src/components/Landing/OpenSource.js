import Link from '@docusaurus/Link';
import React, { useEffect, useState } from 'react';
import classNames from 'classnames';
import { GitBranch, Star } from 'react-feather';

import Heading from './Heading';
import Oval from './Oval';
import styles from './index.module.css';

function Badge({ border, number, text, to, icon: Icon, color }) {
  return (
    <Link to={to} className={classNames(styles.badge)}>
      <img
        className={styles['badge-image']}
        src={`assets/border-0${border}.svg`}
        alt="border-image"
      />
      {Icon && (
        <Oval className={styles['top-right-corner']}>
          <Icon color={color} />
        </Oval>
      )}
      <div className={classNames(styles['badge-centered'], styles.centered)}>
        <div style={{ color }} className={styles['badge-number']}>
          {number}
        </div>
        <p className={styles.text}>{text}</p>
      </div>
    </Link>
  );
}

function cacheKey(key) {
  return `__cached-fetch-${key}`;
}

function cacheGet(key) {
  try {
    const { expires, data } = JSON.parse(localStorage.getItem(cacheKey(key)));
    return { expired: !expires || expires < Date.now(), data };
  } catch (error) {
    return { expired: true, data: null };
  }
}

function cacheSet(key, data, expires = 0) {
  try {
    localStorage.setItem(cacheKey(key), JSON.stringify({ data, expires }));
  } catch (error) {
    // Nothing.
  }
}

function cached(key, fallback, expires) {
  const { expired, data } = cacheGet(key);

  if (expired) {
    return fallback()
      .then(data => {
        cacheSet(key, data, expires);
        return data;
      })
      .catch(() => data || 0);
  }

  return Promise.resolve(data);
}

function formatNumber(number) {
  return number.toLocaleString('en-US');
}

function getGitHubStats() {
  const twoMinutes = 2 * 60 * 1000;
  return cached(
    'github',
    () =>
      fetch('https://api.github.com/repos/autoai-org/aid')
        .then(response => response.json())
        .then(({ stargazers_count: stars, forks_count: forks }) => ({
          stars,
          forks
        })),
    Date.now() + twoMinutes
  );
}

function useStats() {
  const [stars, setStars] = useState(null);
  const [forks, setForks] = useState(null);

  useEffect(() => {
    getGitHubStats().then(({ forks, stars }) => {
      forks && setForks(formatNumber(forks));
      stars && setStars(formatNumber(stars));
    });
  }, [stars, forks]);

  return {
    stars,
    forks,
  };
}

export default function OpenSource() {
  const { stars, forks } = useStats();
  return (
    <div className="section text--center">
      <p
        className={classNames(
          styles.centered,
          styles.text,
          styles['heading-helper']
        )}
      >
      </p>
      <Heading>
        Trusted by Open Source
        <br />
        community
      </Heading>
      <div className={classNames('row', styles.badges)}>
        <Badge
          to="https://github.com/vazco/uniforms/stargazers"
          text="Stars"
          border={1}
          number={stars}
          icon={Star}
          color="#723CFF"
        />
        <Badge
          to="https://github.com/vazco/uniforms/network/members"
          text="Forks"
          border={2}
          number={forks}
          icon={GitBranch}
          color="#3FBBFE"
        />
      </div>
    </div>
  );
}
